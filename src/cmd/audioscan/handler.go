package main

import (
	"database/sql"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	base   *handler.MsgHandler
	config *pb.Configuration
	db     *sql.DB
}

func NewMsgHandler(config *pb.Configuration, db *sql.DB, logger types.Logger, nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base:   handler.NewMsgHandler(logger, nc),
		config: config,
		db:     db,
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_AUDIO_SCAN)),
			Handler: m.handlerAudioScan,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
	if err := m.db.Close(); err != nil {
		m.base.Logger().Error(err)
	}
}
