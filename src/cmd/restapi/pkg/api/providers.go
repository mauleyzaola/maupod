package api

import (
	"errors"
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/rules"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/providers/discogs"
)

func (a *ApiServer) ProviderMetaCoverGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	providerResults, err := discogs.Search(p.r.URL.Query())
	if err != nil {
		status = http.StatusForbidden
		return
	}
	var results []*discogs.Result
	for _, v := range providerResults.Results {
		if v.CoverImage == "" {
			continue
		}
		v.UUID = helpers.NewUUID()
		results = append(results, v)
	}
	if len(results) == 0 {
		results = []*discogs.Result{}
	}

	result = results
	return
}

func (a *ApiServer) ProviderMetaCoverPut(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var input protos.ArtworkDownloadInput
	var output protos.ArtworkDownloadOutput
	if err = p.Decode(&input); err != nil {
		status = http.StatusBadRequest
		return
	}
	input.AlbumIdentifier = p.Param("album_identifier")
	if err = broker.DoRequest(a.nc, protos.Message_MESSAGE_ARTWORK_DOWNLOAD, &input, &output, rules.Timeout(a.config)); err != nil {
		status = http.StatusBadRequest
		return
	}

	if output.Error != "" {
		status = http.StatusBadRequest
		err = errors.New(output.Error)
		return
	}
	result = output
	return
}
