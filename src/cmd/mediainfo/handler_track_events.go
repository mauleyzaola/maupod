package main

import (
	"context"
	"log"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (m *MsgHandler) handlerTrackPlayCountIncrease(msg *nats.Msg) {
	var input protos.TrackPlayedInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		log.Println(err)
		return
	}
	if !checkShaExists(input.Media) {
		return
	}
	ctx := context.Background()
	if err = insertEvent(ctx, m.db, protos.Message_MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE, input.Media); err != nil {
		log.Println(err)
		return
	}
	log.Println("handlerTrackPlayCountIncrease: " + input.Media.Track)
}

func (m *MsgHandler) handlerTrackSkipped(msg *nats.Msg) {
	var input protos.TrackSkippedInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		log.Println(err)
		return
	}
	if !checkShaExists(input.Media) {
		return
	}
	ctx := context.Background()
	if err = insertEvent(ctx, m.db, protos.Message_MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE, input.Media); err != nil {
		log.Println(err)
		return
	}
	log.Println("handlerTrackSkipped: " + input.Media.Track)
}

func insertEvent(ctx context.Context, conn boil.ContextExecutor, event protos.Message, media *protos.Media) error {
	var mediaEvent = orm.MediaEvent{
		ID:    helpers.NewUUID(),
		Sha:   media.Sha,
		TS:    time.Now(),
		Event: int(event),
	}
	return mediaEvent.Insert(ctx, conn, boil.Infer())
}

func checkShaExists(media *protos.Media) bool {
	if media.Sha == "" {
		log.Printf("missing sha for file: %s\n", media.Location)
		return false
	}
	return true
}
