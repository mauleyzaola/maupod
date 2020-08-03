package api

import (
	"net/http"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/paths"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

func (a *ApiServer) AudioScanPost(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input pb.ScanDirectoryAudioFilesInput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}
	// consider the root to be relative to media directory, not absolute
	input.Root = paths.LocationPath(input.Root)

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
