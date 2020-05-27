package main

import (
	"strconv"

	"github.com/mauleyzaola/maupod/src/server/pkg/handler"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	base   *handler.MsgHandler
	config *pb.Configuration
}

func NewMsgHandler(config *pb.Configuration, logger types.Logger, nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base:   handler.NewMsgHandler(logger, nc),
		config: config,
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_INFO)),
			Handler: m.handlerMediaInfo,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
}
