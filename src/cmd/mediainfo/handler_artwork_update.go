package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/boil"
)

func (m *MsgHandler) handlerMediaUpdateArtwork(msg *nats.Msg) {
	var input protos.ArtworkUpdateInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		log.Println(err)
		return
	}
	ctx := context.Background()
	conn := m.db
	store := dbdata.NewMediaStore()

	if err = ArtworkDbUpdate(ctx, conn, input.Media, store); err != nil {
		log.Println(err)
	}
}

func ArtworkDbUpdate(ctx context.Context, conn boil.ContextExecutor, media *protos.Media, store *dbdata.MediaStore) error {
	var cols = orm.MediumColumns
	return store.Update(ctx, conn, media,
		cols.ImageLocation,
		cols.LastImageScan,
		cols.ModifiedDate,
	)
}
