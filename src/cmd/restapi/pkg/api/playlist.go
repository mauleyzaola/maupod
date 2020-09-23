package api

import (
	"errors"
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/mauleyzaola/maupod/src/pkg/pb"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
)

func (a *ApiServer) PlaylistGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	playlist, err := orm.FindPlaylist(p.ctx, p.conn, p.Param("id"))
	if err != nil {
		status = http.StatusNotFound
		return
	}
	result = conversion.PlaylistFromORM(playlist)
	return
}

func (a *ApiServer) PlaylistsGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var filter dbdata.QueryFilter
	if err = p.DecodeQuery(&filter); err != nil {
		status = http.StatusBadRequest
		return
	}
	store := dbdata.NewMediaStore()
	mods := store.PlaylistMods(filter)
	playlists, err := orm.Playlists(mods...).All(p.ctx, p.conn)
	if err != nil {
		status = http.StatusBadRequest
		return
	}
	result = conversion.PlaylistsFromORM(playlists...)
	return
}

func (a *ApiServer) PlaylistPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var playlist pb.PlayList
	if err = p.Decode(&playlist); err != nil {
		status = http.StatusBadRequest
		return
	}
	if playlist.Name == "" {
		status = http.StatusBadRequest
		err = errors.New("missing name")
		return
	}
	playlist.Id = helpers.NewUUID()
	v := conversion.PlaylistToORM(&playlist)
	if err = v.Insert(p.ctx, p.conn, boil.Infer()); err != nil {
		status = http.StatusInternalServerError
		return
	}
	result = playlist
	status = http.StatusCreated
	return
}

func (a *ApiServer) PlaylistPut(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var playlist pb.PlayList
	if err = p.Decode(&playlist); err != nil {
		status = http.StatusBadRequest
		return
	}
	playlist.Id = p.Param("id")
	v := conversion.PlaylistToORM(&playlist)
	var cols = orm.PlaylistColumns
	rowCount, err := v.Update(p.ctx, p.conn, boil.Whitelist(cols.Name))
	if err != nil {
		status = http.StatusBadRequest
		return
	}
	if rowCount == 0 {
		status = http.StatusNotFound
		return
	} else if rowCount != 1 {
		status = http.StatusConflict
		return
	}
	status = http.StatusAccepted
	return
}

func (a *ApiServer) PlaylistDelete(p TransactionExecutorParams) (status int, result interface{}, err error) {
	return
}

func (a *ApiServer) PlaylistItemPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	return
}

func (a *ApiServer) PlaylistItemDelete(p TransactionExecutorParams) (status int, result interface{}, err error) {
	return
}

func (a *ApiServer) PlaylistItemPut(p TransactionExecutorParams) (status int, result interface{}, err error) {
	return
}

func (a *ApiServer) PlaylistItems(p TransactionExecutorParams) (status int, result interface{}, err error) {
	return
}
