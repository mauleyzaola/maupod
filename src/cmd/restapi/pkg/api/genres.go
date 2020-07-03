package api

import (
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *ApiServer) GenresListGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	if result, err = orm.ViewGenres().All(p.ctx, p.conn); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}

// GenreArtworksGet returns available artwork for genres
func (a *ApiServer) GenreArtworksGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var mods []qm.QueryMod
	var cols = orm.MediumColumns
	mods = append(mods, qm.Where(cols.ShaImage+" <> '' "))
	rows, err := orm.Media(mods...).All(p.ctx, p.conn)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	var keys = make(map[string][]string)
	for _, r := range rows {
		genre := r.Genre
		if genre == "" {
			continue
		}
		val := keys[genre]
		val = append(val, r.ShaImage)
		keys[genre] = val
	}

	for k, v := range keys {
		keys[k] = helpers.StringSlice(v).Distinct()
	}

	result = keys
	return
}
