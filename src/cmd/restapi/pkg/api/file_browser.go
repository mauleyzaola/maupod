package api

import (
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
)

func (a *ApiServer) DirectoryReadGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.DirectoryReadInput

	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}

	if result, err = broker.RequestFileBrowserDirectory(a.nc, &input, rules.Timeout(a.config)); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}
