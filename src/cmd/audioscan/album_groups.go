package main

import (
	"context"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AlbumGroupDetection will calculate based on the current media, the album_identifier,
// and whether it is or not a compilation
func AlbumGroupDetection(ctx context.Context, conn boil.ContextExecutor, media *protos.Media) (albumIdentifier string, err error) {
	var mods []qm.QueryMod
	var where = orm.MediumWhere
	var cols = orm.MediumColumns

	// read rows same directory
	mods = append(mods, where.Directory.EQ(media.Directory))

	rows, err := orm.Media(mods...).All(ctx, conn)
	if err != nil {
		return
	}

	// TODO: implement compilation
	for _, v := range rows {
		if v.AlbumIdentifier != "" {
			albumIdentifier = v.AlbumIdentifier
			break
		}
	}
	// not assigned yet, then create one
	if albumIdentifier == "" {
		albumIdentifier = helpers.NewUUID()
	}

	for _, v := range rows {
		if v.AlbumIdentifier != albumIdentifier {
			v.AlbumIdentifier = albumIdentifier
			if _, err = v.Update(ctx, conn, boil.Whitelist(cols.AlbumIdentifier)); err != nil {
				return
			}
		}
	}

	return
}
