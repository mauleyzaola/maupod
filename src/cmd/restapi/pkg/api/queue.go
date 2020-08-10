package api

import (
	"net/http"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
)

func (a *ApiServer) QueueGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.QueueInput
	result, err = broker.RequestQueueList(a.nc, &input, rules.Timeout(a.config))
	if err != nil {
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
	output, err := broker.RequestQueueAdd(a.nc, &input, rules.Timeout(a.config))
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	result = output.Rows
	return
}

func (a *ApiServer) QueueDelete(p TransactionExecutorParams) (status int, result interface{}, err error) {
	val, err := strconv.Atoi(p.Param("index"))
	if err != nil {
		status = http.StatusBadRequest
		return
	}
	var input = pb.QueueInput{
		Index: int64(val),
	}
	output, err := broker.RequestQueueRemove(a.nc, &input, rules.Timeout(a.config))
	if err != nil {
		status = http.StatusBadRequest
		return
	}
	result = output
	return
}
