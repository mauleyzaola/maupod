package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	var input pb.PlaylistDeleteInput
	if err = p.DecodeQuery(&input); err != nil {
		status = http.StatusBadRequest
		return
	}
	input.Id = p.Param("id")

	v := orm.Playlist{ID: input.Id}

	if input.ForceDeleteChildren {
		if _, err = v.PlaylistItems().DeleteAll(p.ctx, p.conn); err != nil {
			status = http.StatusInternalServerError
			return
		}
	}
	rowCount, err := v.Delete(p.ctx, p.conn)
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
	status = http.StatusNoContent
	return
}

func (a *ApiServer) PlaylistItemPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var item pb.PlaylistItem
	if err = p.Decode(&item); err != nil {
		status = http.StatusBadRequest
		return
	}
	item.Playlist = &pb.PlayList{Id: p.Param("id")}
	if item.Media == nil {
		err = errors.New("missing media in request")
		status = http.StatusBadRequest
		return
	}
	// count the play list items
	var where = orm.PlaylistItemWhere
	rowCount, err := orm.PlaylistItems(where.PlaylistID.EQ(item.Playlist.Id)).Count(p.ctx, p.conn)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	// set the new item position to the length of the items for this playlist
	item.Position = int32(rowCount)

	item.Id = helpers.NewUUID()
	v := conversion.PlaylistItemToORM(&item)
	if err = v.Insert(p.ctx, p.conn, boil.Infer()); err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusCreated
	return
}

func (a *ApiServer) PlaylistItemDelete(p TransactionExecutorParams) (status int, result interface{}, err error) {
	id := p.Param("id")
	position, err := strconv.Atoi(p.Param("position"))
	if err != nil {
		status = http.StatusBadRequest
		return
	}
	// delete the provided item in the playlist
	var where = orm.PlaylistItemWhere
	rowCount, err := orm.PlaylistItems(where.PlaylistID.EQ(id), where.Position.EQ(position)).DeleteAll(p.ctx, p.conn)
	if err != nil {
		status = http.StatusBadRequest
		return
	} else if rowCount == 0 {
		status = http.StatusBadRequest
		err = errors.New("wrong position or wrong playlist provided")
		return
	} else if rowCount != 1 {
		status = http.StatusInternalServerError
		err = fmt.Errorf("expected to affect 1 row, but would affect: %d rows instead", rowCount)
		return
	}

	// update the position of the items after the one we deleted
	var cols = orm.PlaylistItemColumns
	nextItems, err := orm.PlaylistItems(where.PlaylistID.EQ(id), where.Position.GT(position)).All(p.ctx, p.conn)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	for _, item := range nextItems {
		item.Position--
		if _, err = item.Update(p.ctx, p.conn, boil.Whitelist(cols.Position)); err != nil {
			status = http.StatusInternalServerError
			return
		}
	}

	status = http.StatusAccepted
	return
}

func (a *ApiServer) PlaylistItemPut(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.PlaylistItem
	var id = p.Param("id")
	positionStr := p.Param("position")
	position, err := strconv.Atoi(positionStr)
	if err != nil {
		err = fmt.Errorf("cannot parse position to int: %s", positionStr)
		status = http.StatusBadRequest
		return
	}
	if position < 0 {
		err = errors.New("position should be equal or greater than zero")
		status = http.StatusBadRequest
		return
	}

	// decode payload
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}

	// basic health check
	if input.Media == nil || input.Media.Id == "" {
		err = errors.New("missing media or media id")
		status = http.StatusBadRequest
		return
	}

	// get current playlist items
	items, err := playlistItemsList(p.ctx, p.conn, id)
	if err != nil {
		status = http.StatusBadRequest
		return
	}
	// check position out of bounds
	if length := len(items) - 1; position > length {
		err = fmt.Errorf("invalid position: %d, max position is: %d", position, length)
		status = http.StatusBadRequest
		return
	}
	// find a matching item with input provided
	var matchedItem *pb.PlaylistItem
	for _, v := range items {
		if v.Media.Id == input.Media.Id && v.Position == input.Position {
			matchedItem = v
			break
		}
	}

	if matchedItem == nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("didn't find any item in playlist at position: %d with media.id: %s", input.Position, input.Media.Id)
			status = http.StatusNotFound
		} else {
			status = http.StatusBadRequest
		}
		return
	}
	if int(matchedItem.Position) == position {
		status = http.StatusBadRequest
		err = errors.New("the position and new position are the same value")
		return
	}

	start := int(matchedItem.Position)
	for i := start; i >= position; i-- {
		items[i] = items[i-1]
	}
	items[position] = matchedItem

	for _, v := range items {
		log.Println("position,track: ", v.Position, v.Media.Track)
	}
	err = fmt.Errorf("xxx")
	return

	//var cols = orm.PlaylistItemColumns
	//for i, v := range clonedItems {
	//	if v.Position == i {
	//		continue
	//	}
	//	v.Position = i
	//	if _, err = v.Update(p.ctx, p.conn, boil.Whitelist(cols.Position)); err != nil {
	//		status = http.StatusInternalServerError
	//		return
	//	}
	//}

	return
}

func (a *ApiServer) PlaylistItems(p TransactionExecutorParams) (status int, result interface{}, err error) {
	if result, err = playlistItemsList(p.ctx, p.conn, p.Param("id")); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}

func playlistItemsList(ctx context.Context, conn boil.ContextExecutor, playlistID string) ([]*pb.PlaylistItem, error) {
	var mods []qm.QueryMod
	var cols = orm.PlaylistItemColumns
	var where = orm.PlaylistItemWhere
	mods = append(mods, where.PlaylistID.EQ(playlistID))
	mods = append(mods, qm.OrderBy(cols.Position+" asc"))
	items, err := orm.PlaylistItems(mods...).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	var result []*pb.PlaylistItem
	for i, v := range items {
		var media *orm.Medium
		item := conversion.PlaylistItemFromORM(v)
		item.Position = int32(i)
		media, err = orm.FindMedium(ctx, conn, item.Media.Id)
		if err != nil {
			return nil, err
		}
		item.Media = conversion.MediaFromORM(media)
		result = append(result, item)
	}
	return result, nil
}
