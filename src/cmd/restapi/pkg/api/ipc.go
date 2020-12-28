package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/protos"
)

func (a *ApiServer) IPCPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input protos.IPCInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}

	if err = broker.RequestIPCCommand(a.nc, &input, time.Second*time.Duration(a.config.Delay)); err != nil {
		status = http.StatusBadRequest
		return
	}

	return
}

func (a *ApiServer) VolumeChange(p TransactionExecutorParams) (status int, result interface{}, err error) {
	payload := &struct {
		Offset int64
	}{}
	if err = p.Decode(payload); err != nil {
		status = http.StatusBadRequest
		return
	}
	input := &protos.VolumeChangeInput{Offset: int32(payload.Offset)}

	data, err := helpers.ProtoMarshalJSON(input)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	subject := strconv.Itoa(int(protos.Message_MESSAGE_VOLUME_CHANGE))
	msg, err := a.nc.Request(subject, data, a.Timeout())
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	output := &protos.VolumeChangeOutput{}
	if err = helpers.ProtoUnmarshal(msg.Data, output); err != nil {
		status = http.StatusInternalServerError
		return
	}

	if !output.Ok {
		status = http.StatusBadRequest
		err = fmt.Errorf(output.Error)
		return
	}
	result = &output
	return
}
