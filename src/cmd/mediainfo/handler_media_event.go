package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (m *MsgHandler) handlerMediaEventUpsert(msg *nats.Msg) {
	var input pb.MediaEventInput

	if err := helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	mediaEvent := &orm.MediaEvent{
		ID:    input.Id,
		Sha:   input.Sha,
		TS:    helpers.TsToTime2(input.Ts),
		Event: int(input.Event),
	}
	ctx := context.Background()
	conn := m.db
	if err := mediaEvent.Insert(ctx, conn, boil.Infer()); err != nil {
		log.Println(err)
	}
	log.Println("[INFO] successfully imported event")
}
