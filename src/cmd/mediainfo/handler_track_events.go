package main

import (
	"context"
	"log"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (m *MsgHandler) handlerTrackPlayCountIncrease(msg *nats.Msg) {
	var input pb.TrackPlayedInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	if !checkShaExists(input.Media) {
		return
	}
	ctx := context.Background()
	if err = insertEvent(ctx, m.db, pb.Message_MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE, input.Media); err != nil {
		m.base.Logger().Error(err)
		return
	}
	m.base.Logger().Info("handlerTrackPlayCountIncrease: " + input.Media.Track)
}

func (m *MsgHandler) handlerTrackSkipped(msg *nats.Msg) {
	var input pb.TrackSkippedInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	if !checkShaExists(input.Media) {
		return
	}
	ctx := context.Background()
	if err = insertEvent(ctx, m.db, pb.Message_MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE, input.Media); err != nil {
		m.base.Logger().Error(err)
		return
	}
	m.base.Logger().Info("handlerTrackSkipped: " + input.Media.Track)
}

func insertEvent(ctx context.Context, conn boil.ContextExecutor, event pb.Message, media *pb.Media) error {
	var mediaEvent = orm.MediaEvent{
		ID:    helpers.NewUUID(),
		Sha:   media.Sha,
		TS:    time.Now(),
		Event: int(event),
	}
	return mediaEvent.Insert(ctx, conn, boil.Infer())
}

func checkShaExists(media *pb.Media) bool {
	if media.Sha == "" {
		log.Printf("missing sha for file: %s\n", media.Location)
		return false
	}
	return true
}
