package main

import (
	"strconv"

	"github.com/mauleyzaola/maupod/src/server/cmd/player/pkg"
	"github.com/mauleyzaola/maupod/src/server/pkg/handler"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
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
	m.base.Close()
}

// InitializeIPC is required so we can be sure the first filename we receive we initialize the ipc object
func (m *MsgHandler) InitializeIPC(filename string) error {
	if m.isInitialized {
		return nil
	}
	processor, err := pkg.NewMpvProcessor(filename)
	if err != nil {
		return err
	}
	if m.ipc, err = pkg.NewIPC(processor); err != nil {
		return err
	}
	m.isInitialized = true
	return nil
}
