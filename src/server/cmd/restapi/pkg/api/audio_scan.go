package api

import (
	"net/http"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/broker"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func (a *ApiServer) AudioScanPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.ScanDirectoryAudioFilesInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}
	// we need to set from the caller, when the file was requested to be scanned
	// so this value goes to db
	input.ScanDate = helpers.TimeToTs2(time.Now())

	if err = broker.PublishBroker(a.nc, pb.Message_MESSAGE_AUDIO_SCAN, &input); err != nil {
		status = http.StatusInternalServerError
		return
	}
	return
}

func (a *ApiServer) ArtworkScanPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	return
}
