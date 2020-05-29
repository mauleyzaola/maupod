package api

import (
	"net/http"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
)

func (a *ApiServer) PerformersGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var filter data.MediaFilter
	if err = p.DecodeQuery(&filter); err != nil {
		status = http.StatusBadRequest
		return
	}
	filter.Distinct = p.Param("field")
	if result, err = a.dm.DistinctList(p.ctx, p.conn, filter); err != nil {
		status = http.StatusBadRequest
		return
	}
	return
}
