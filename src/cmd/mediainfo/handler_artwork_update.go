package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/boil"
)

func (m *MsgHandler) handlerMediaUpdateArtwork(msg *nats.Msg) {
	var input pb.ArtworkUpdateInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		log.Println(err)
		return
	}
	ctx := context.Background()
	conn := m.db
	store := &dbdata.MediaStore{}

	if err = ArtworkDbUpdate(ctx, conn, input.Media, store); err != nil {
		log.Println(err)
	}
}

func ArtworkDbUpdate(ctx context.Context, conn boil.ContextExecutor, media *pb.Media, store *dbdata.MediaStore) error {
	var cols = orm.MediumColumns
	return store.Update(ctx, conn, media,
		cols.ImageLocation,
		cols.LastImageScan,
		cols.ModifiedDate,
	)
}
