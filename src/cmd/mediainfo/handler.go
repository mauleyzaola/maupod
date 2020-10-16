package main

import (
	"database/sql"
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

func NewMsgHandler(config *pb.Configuration, nc *nats.Conn, db *sql.DB) *MsgHandler {
	return &MsgHandler{
		base:   handler.NewMsgHandler(nc),
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
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_UPDATE_SHA)),
			Handler: m.handlerUpdateSHA,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_QUEUE_LIST)),
			Handler: m.handlerQueueList,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_QUEUE_ADD)),
			Handler: m.handlerQueueAdd,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_QUEUE_REMOVE)),
			Handler: m.handlerQueueRemove,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_DIRECTORY_READ)),
			Handler: m.handlerReadDirectory,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_SPECTRUM_GENERATE)),
			Handler: m.handlerMediaSpectrumGenerate,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_DB_SELECT)),
			Handler: m.handlerMediaInfoDBSelect,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_MICRO_SERVICE_MEDIAINFO)),
			Handler: m.handlerMicroService,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_UPSERT_MEDIA_EVENT)),
			Handler: m.handlerMediaEventUpsert,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
}
