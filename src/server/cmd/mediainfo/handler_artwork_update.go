package main

import (
	"context"

	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/boil"
)

func (m *MsgHandler) handlerMediaUpdateArtwork(msg *nats.Msg) {
	var input pb.ArtworkUpdateInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	ctx := context.Background()
	conn := m.db
	store := &dbdata.MediaStore{}

	if err = ArtworkDbUpdate(ctx, conn, input.Media, store); err != nil {
		m.base.Logger().Error(err)
	}
}

func ArtworkDbUpdate(ctx context.Context, conn boil.ContextExecutor, media *pb.Media, store *dbdata.MediaStore) error {
	var cols = orm.MediumColumns
	return store.Update(ctx, conn, media,
		cols.ImageLocation,
		cols.ShaImage,
		cols.LastImageScan,
		cols.ModifiedDate,
	)
}