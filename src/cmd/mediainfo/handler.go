package main

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/handler"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	base       *handler.MsgHandler
	config     *pb.Configuration
	db         *sql.DB
	queueItems types.Medias
}

func NewMsgHandler(config *pb.Configuration, logger types.Logger, nc *nats.Conn, db *sql.DB) *MsgHandler {
	queueItems, err := queueList(context.Background(), db)
	if err != nil {
		log.Println("[ERROR] queueList() ", err)
	}

	return &MsgHandler{
		base:       handler.NewMsgHandler(logger, nc),
		config:     config,
		db:         db,
		queueItems: queueItems,
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
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_QUEUE_SAVE)),
			Handler: m.handlerQueueSave,
		},
		handler.Subscription{
			Subject: strconv.Itoa(int(pb.Message_MESSAGE_DIRECTORY_READ)),
			Handler: m.handlerReadDirectory,
		},
	)
}

func (m *MsgHandler) Close() {
	m.base.Close()
}
