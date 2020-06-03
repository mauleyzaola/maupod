package main

import (
	"context"

	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// AlbumGroupDetection will calculate based on the current media, the album_identifier,
// and whether it is or not a compilation
func AlbumGroupDetection(ctx context.Context, conn boil.ContextExecutor, media *pb.Media) (albumIdentifier string, isCompilation bool, err error) {
	var mods []qm.QueryMod
	var where = orm.MediumWhere
	var cols = orm.MediumColumns

	// read rows same album and performer
	mods = append(mods, where.Performer.EQ(media.Performer))
	mods = append(mods, where.Album.EQ(media.Album))

	rows, err := orm.Media(mods...).All(ctx, conn)
	if err != nil {
		return
	}
	if len(rows) == 0 {
		// there is no way of knowing what this is yet, we assume it is a compilation
		isCompilation = true
		return
	}
	isCompilation = false
	for _, v := range rows {
		if v.AlbumIdentifier != "" {
			albumIdentifier = v.AlbumIdentifier
			break
		}
	}
	// not assigned yet, then create one
	albumIdentifier = helpers.NewUUID()

	for _, v := range rows {
		if v.AlbumIdentifier != albumIdentifier || v.IsCompilation != isCompilation {
			v.AlbumIdentifier = albumIdentifier
			v.IsCompilation = isCompilation
			if _, err = v.Update(ctx, conn, boil.Whitelist(cols.AlbumIdentifier, cols.IsCompilation)); err != nil {
				return
			}
		}
	}

	return
}
