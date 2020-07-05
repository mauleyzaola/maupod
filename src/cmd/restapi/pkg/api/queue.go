package api

import (
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *ApiServer) QueueListGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	rows, err := orm.MediaQueues(qm.OrderBy(orm.MediaQueueColumns.ID+" asc")).All(p.ctx, p.conn)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	var ids []interface{}
	for _, v := range rows {
		ids = append(ids, v.MediaID)
	}

	medias, err := orm.Media(qm.WhereIn(orm.MediumColumns.ID+" in ?", ids...)).All(p.ctx, p.conn)
	if err != nil {
		status = http.StatusInternalServerError
		return
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
	result = res
	return
}

func (a *ApiServer) QueuePost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.QueueInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}
	return
}
