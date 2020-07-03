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

func NewMsgHandler(config *pb.Configuration, logger types.Logger, nc *nats.Conn, db *sql.DB) *MsgHandler {
	return &MsgHandler{
		base:   handler.NewMsgHandler(logger, nc),
		config: config,
		db:     db,
	}
}

func (m *MsgHandler) Register() error {
	return m.base.Register(
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_INFO)),
			Handler: m.handlerMediaInfo,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_UPDATE_ARTWORK)),
			Handler: m.handlerMediaUpdateArtwork,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_UPDATE)),
			Handler: m.handlerMediaUpdateDb,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE)),
			Handler: m.handlerTrackPlayCountIncrease,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE)),
			Handler: m.handlerTrackSkipped,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
}