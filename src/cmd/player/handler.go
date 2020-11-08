package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/mauleyzaola/maupod/src/cmd/player/pkg"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type MsgHandler struct {
	base          *handler.MsgHandler
	ipc           *pkg.IPC
	isInitialized bool
	rc            *redis.Client
	timeout       time.Duration

	// state of last media played
	lastMediaPlayed *protos.LastPlayedMediaInput
}

func NewMsgHandler(nc *nats.Conn, rc *redis.Client, timeout time.Duration) *MsgHandler {
	return &MsgHandler{
		base:          handler.NewMsgHandler(nc),
		isInitialized: false,
		rc:            rc,
		timeout:       timeout,
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_IPC)),
			Handler: m.handlerIPC,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_SOCKET_TRACK_POSITION_PERCENT_CHANGE)),
			Handler: m.handlerPositionPercentChange,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_MICRO_SERVICE_PLAYER)),
			Handler: m.handlerMicroService,
		},
	)
}

func (m *MsgHandler) Close() {
	// store the last position played and the last media played
	// save this information to redis
	ctx := context.Background()
	key := strconv.Itoa(int(protos.Message_IPC_LAST_PLAYED_MEDIA))
	if m.lastMediaPlayed != nil {
		data, err := helpers.ProtoMarshal(m.lastMediaPlayed)
		if err != nil {
			log.Println(err)
		}
		if err := m.rc.Set(ctx, key, data, 0).Err(); err != nil {
			log.Println(err)
		}
	} else {
		// TODO: enable this when completed debugging
		// clear last media played for next player start up
		//if err := m.rc.Del(ctx, key).Err(); err != nil {
		//	log.Println(err)
		//}
	}

	log.Println("saving state of player to redis")
	if m.isInitialized {
		if err := m.ipc.Terminate(); err != nil {
			log.Println(err)
		}
	}
	m.base.Close()
}

// InitializeIPC is required so we can be sure the first filename we receive we initialize the ipc object
func (m *MsgHandler) InitializeIPC(filename string) error {
	if m.isInitialized {
		return nil
	}
	processor, err := pkg.NewMpvProcessor()
	if err != nil {
		return err
	}
	var publishFn broker.PublisherFunc = func(subject protos.Message, payload proto.Message) error {
		return broker.PublishBroker(m.base.NATS(), subject, payload)
	}
	var requestFn broker.RequestFunc = func(subject protos.Message, input, output proto.Message) error {
		// TODO: timeout should come from configuration
		return broker.DoRequest(m.base.NATS(), subject, input, output, time.Second)
	}
	var publishFnJSON broker.PublisherFuncJSON = func(subject protos.Message, payload interface{}) error {
		val, ok := payload.(proto.Message)
		if !ok {
			log.Println("[ERROR] cannot cast to protos.Message: ", payload)
		}
		return broker.PublishBrokerJSON(m.base.NATS(), subject, val)
	}
	control := pkg.NewPlayerControl(publishFn, publishFnJSON, requestFn)

	// pass a function handler that stores the position played and the media, so we can continue playing on re-starting maupod-player
	var lastPlayedState = func(lastMedia *protos.Media, percent float64) {
		// TODO: probably need to handle this on another trigger like STOP?
		if lastMedia == nil {
			m.lastMediaPlayed = nil
			return
		}
		if m.lastMediaPlayed == nil {
			m.lastMediaPlayed = &protos.LastPlayedMediaInput{}
		}
		m.lastMediaPlayed.Media = lastMedia
		m.lastMediaPlayed.Percent = percent
	}
	if m.ipc, err = pkg.NewIPC(processor, control, lastPlayedState); err != nil {
		return err
	}
	m.isInitialized = true
	return nil
}

func (m *MsgHandler) Start() error {
	// check if there is a pending media from last time to be resumed
	ctx := context.Background()
	cmd := m.rc.Get(ctx, strconv.Itoa(int(protos.Message_IPC_LAST_PLAYED_MEDIA)))
	if err := cmd.Err(); err != nil {
		log.Println("no pending media to be resumed")
		return nil
	}
	var resumedMedia protos.LastPlayedMediaInput
	data, err := cmd.Bytes()
	if err != nil {
		log.Println(err)
		return err
	}
	if err = helpers.ProtoUnmarshal(data, &resumedMedia); err != nil {
		log.Println(err)
		return err
	}

	media := resumedMedia.Media
	var filename string
	if val := media.Location; val != "" {
		var location = paths.LocationPath(val)
		filename = paths.MediaFullPathAudioFile(location)
		if err := m.InitializeIPC(filename); err != nil {
			log.Println(err)
			return err
		}
	}
	media.Location = filename
	if err := m.InitializeIPC(filename); err != nil {
		log.Println(err)
		return err
	}

	log.Printf("[INFO] found a tracks that needs to be resumed: %s at: %v%%\n", resumedMedia.Media.Track, resumedMedia.Percent)
	var ipcInputs = []*protos.IPCInput{
		{
			Media:   resumedMedia.Media,
			Value:   "",
			Command: protos.Message_IPC_PLAY,
		},
		{
			Media:   resumedMedia.Media,
			Value:   "",
			Command: protos.Message_IPC_PAUSE,
		},
	}
	for _, ipcInput := range ipcInputs {
		if err = broker.RequestIPCCommand(m.base.NATS(), ipcInput, m.timeout); err != nil {
			log.Println(err)
		}
	}
	posInput := &protos.SocketTrackPositionChangeInput{
		Media:   media,
		Percent: resumedMedia.Percent,
	}
	//	 TODO: move this function to helpers package
	// need to send this message as json so it matches the signature from socket server
	jp := runtime.JSONPb{}
	if data, err = jp.Marshal(posInput); err != nil {
		log.Println(err)
		return err
	}
	// TODO: this should be a signal from the handler instead of fixed duration
	time.Sleep(time.Second * 2)
	if err = m.base.NATS().Publish(strconv.Itoa(int(protos.Message_MESSAGE_SOCKET_TRACK_POSITION_PERCENT_CHANGE)), data); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
