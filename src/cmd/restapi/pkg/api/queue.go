package api

import (
	"context"
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func queueList(ctx context.Context, conn boil.ContextExecutor) (types.Medias, error) {
	rows, err := orm.MediaQueues(qm.OrderBy(orm.MediaQueueColumns.ID+" asc")).All(ctx, conn)
	if err != nil {
		return nil, err
	}

	var ids []interface{}
	for _, v := range rows {
		ids = append(ids, v.MediaID)
	}

	medias, err := orm.Media(qm.WhereIn(orm.MediumColumns.ID+" in ?", ids...)).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	var keys = make(map[string]*pb.Media)
	for _, v := range medias {
		keys[v.ID] = conversion.MediaFromORM(v)
	}

	var res []*pb.Media
	for _, v := range rows {
		val, ok := keys[v.MediaID]
		if !ok {
			continue
		}
		res = append(res, val)
	}
	return res, nil
}

func (a *ApiServer) QueueGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	if result, err = queueList(p.ctx, p.conn); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}

func (a *ApiServer) QueuePost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.QueueInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}

	list, err := queueList(p.ctx, p.conn)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	if list, err = list.InsertAt(input.Media, int(input.Index)); err != nil {
		status = http.StatusBadRequest
		return
	}
	result = list
	return
}

func (a *ApiServer) QueueDelete(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input = struct {
		Index int `schema:"index"`
	}{}
	if err = p.DecodeQuery(&input); err != nil {
		status = http.StatusBadRequest
		return
	}

	list, err := queueList(p.ctx, p.conn)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	if list, err = list.RemoveAt(input.Index); err != nil {
		status = http.StatusBadRequest
		return
	}
	result = list
	return
}
