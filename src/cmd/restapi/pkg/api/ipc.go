package api

import (
	"net/http"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

func (a *ApiServer) IPCPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.IPCInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}

	if result, err = broker.RequestIPCCommand(a.nc, &input, time.Second*time.Duration(a.config.Delay)); err != nil {
		status = http.StatusInternalServerError
		return
	}

	return
}
