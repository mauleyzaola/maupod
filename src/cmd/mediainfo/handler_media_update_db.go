package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerMediaUpdateDb(msg *nats.Msg) {
	var err error
	var input pb.MediaInfoInput

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	if input.Media == nil {
		log.Println("[ERROR] missing media parameter")
		return
	}
	if input.Media.Id == "" {
		log.Println("[ERROR] cannot update media row: missing media.ID")
		return
	}

	ctx := context.Background()
	conn := m.db
	store := dbdata.MediaStore{}
	input.Media.ModifiedDate = helpers.TimeToTs(helpers.Now())
	var cols = orm.MediumColumns
	var fields = []string{cols.TrackNameTotal, cols.Track, cols.TrackPosition, cols.Album, cols.Comment, cols.Genre, cols.Performer, cols.ModifiedDate}
	if err = store.Update(ctx, conn, input.Media, fields...); err != nil {
		log.Println(err)
		return
	}

	return

}
