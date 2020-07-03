package main

import (
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	base *handler.MsgHandler
}

func NewMsgHandler(logger types.Logger, nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base: handler.NewMsgHandler(logger, nc),
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_REST_API_READY)),
			Handler: m.handlerPing,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
}
