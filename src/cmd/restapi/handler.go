package main

import (
	"strconv"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	base *handler.MsgHandler
}

func NewMsgHandler(nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base: handler.NewMsgHandler(nc),
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_MICRO_SERVICE_RESTAPI)),
			Handler: m.handlerMicroService,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
}
