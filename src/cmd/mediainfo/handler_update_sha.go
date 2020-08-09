package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (m *MsgHandler) handlerUpdateSHA(msg *nats.Msg) {
	var input pb.MediaUpdateSHAInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		log.Println(err)
		return
	}
	ctx := context.Background()
	conn := m.db
	var mods []qm.QueryMod
	var where = orm.MediaEventWhere
	var cols = orm.MediaEventColumns
	mods = append(mods, where.Sha.EQ(input.OldSHA))
	events, err := orm.MediaEvents(mods...).All(ctx, conn)
	if err != nil {
		log.Println(err)
		return
	}
	for _, v := range events {
		v.Sha = input.NewSHA
		if _, err = v.Update(ctx, conn, boil.Whitelist(cols.Sha)); err != nil {
			log.Println(err)
			return
		}
	}
}
