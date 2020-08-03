package main

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/cmd/player/pkg"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type MsgHandler struct {
	base          *handler.MsgHandler
	ipc           *pkg.IPC
	isInitialized bool
}

func NewMsgHandler(logger types.Logger, nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base:          handler.NewMsgHandler(logger, nc),
		isInitialized: false,
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_IPC)),
			Handler: m.handlerIPC,
		},
	)
}

func (m *MsgHandler) Close() {
	if m.isInitialized {
		if err := m.ipc.Terminate(); err != nil {
			m.base.Logger().Error(err)
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
	var publishFn broker.PublisherFunc = func(subject pb.Message, payload proto.Message) error {
		return broker.PublishBroker(m.base.NATS(), subject, payload)
	}
	var requestFn broker.RequestFunc = func(subject pb.Message, input, output proto.Message) error {
		// TODO: timeout should come from configuration
		return broker.DoRequest(m.base.NATS(), subject, input, output, time.Second)
	}
	var publishFnJSON broker.PublisherFuncJSON = func(subject pb.Message, payload interface{}) error {
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		return m.base.NATS().Publish(strconv.Itoa(int(subject)), data)
	}
	control := pkg.NewPlayerControl(publishFn, publishFnJSON, requestFn)
	if m.ipc, err = pkg.NewIPC(processor, control); err != nil {
		return err
	}
	m.isInitialized = true
	return nil
}
