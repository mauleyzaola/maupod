package api

import (
	"net/http"

	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
)

func (a *ApiServer) GenresListGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	if result, err = orm.ViewGenres().All(p.ctx, p.conn); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}
