package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	base   *handler.MsgHandler
	db     *sql.DB
	config *protos.Configuration
}

func NewMsgHandler(db *sql.DB, config *protos.Configuration, nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		base:   handler.NewMsgHandler(nc),
		db:     db,
		config: config,
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_ARTWORK_SCAN)),
			Handler: m.handlerArtworkExtract,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_TAG_UPDATE)),
			Handler: m.handlerTagUpdate,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_MICRO_SERVICE_ARTWORK)),
			Handler: m.handlerMicroService,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(protos.Message_MESSAGE_ARTWORK_DOWNLOAD)),
			Handler: m.handlerArtworkDownload,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
	if err := m.db.Close(); err != nil {
		log.Println(err)
	}
}
