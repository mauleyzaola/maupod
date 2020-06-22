package main

import (
	"context"

	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerMediaUpdateDb(msg *nats.Msg) {
	var err error
	var input pb.MediaInfoInput

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		return
	}

	ctx := context.Background()
	conn := m.db
	store := dbdata.MediaStore{}
	input.Media.ModifiedDate = helpers.TimeToTs(helpers.Now())
	var cols = orm.MediumColumns
	var fields = []string{cols.TrackNameTotal, cols.Track, cols.TrackPosition, cols.Album, cols.Track, cols.Comment, cols.Genre, cols.Performer, cols.ModifiedDate}
	if err = store.Update(ctx, conn, input.Media, fields...); err != nil {
		m.base.Logger().Error(err)
		return
	}

	return

}
