package api

import (
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/broker"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func (a *ApiServer) AudioScanPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.ScanDirectoryAudioFilesInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}

	data, err := proto.Marshal(&input)
	if err = broker.PublishMessage(a.nc, strconv.Itoa(int(pb.Message_MESSAGE_AUDIO_SCAN)), data); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}

func (a *ApiServer) ArtworkScanPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	return
}
