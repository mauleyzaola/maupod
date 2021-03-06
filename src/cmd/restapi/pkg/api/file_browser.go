package api

import (
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/protos"
)

func (a *ApiServer) DirectoryReadGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input protos.DirectoryReadInput

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

func (a *ApiServer) FileSync(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input protos.SyncFilesInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}
	input.TargetDirectory = paths.SyncRootDirectory()
	nc := a.nc
	if err = broker.PublishBroker(nc, protos.Message_MESSAGE_SYNC_FILES, &input); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}
