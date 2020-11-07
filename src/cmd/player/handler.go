package main

import (
	"log"
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/cmd/player/pkg"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type MsgHandler struct {
	base          *handler.MsgHandler
	ipc           *pkg.IPC
	isInitialized bool
}

func NewMsgHandler(nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base:          handler.NewMsgHandler(nc),
		isInitialized: false,
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
	if m.ipc, err = pkg.NewIPC(processor, control); err != nil {
		return err
	}
	m.isInitialized = true
	return nil
}
