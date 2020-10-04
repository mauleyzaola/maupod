package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	base   *handler.MsgHandler
	config *pb.Configuration
	db     *sql.DB
}

func NewMsgHandler(config *pb.Configuration, db *sql.DB, nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base:   handler.NewMsgHandler(nc),
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
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MICRO_SERVICE_AUDIOSCAN)),
			Handler: m.handlerMicroService,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_SYNC_FILES)),
			Handler: m.handlerSyncFiles,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_SHA_SCAN)),
			Handler: m.handlerSHAScan,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
	if err := m.db.Close(); err != nil {
		log.Println(err)
	}
}
