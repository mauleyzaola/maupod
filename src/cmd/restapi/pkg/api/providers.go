package api

import (
	"net/http"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/providers/discogs"
)

func (a *ApiServer) ProviderMetaCoverGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	providerResults, err := discogs.Search(p.r.URL.Query())
	if err != nil {
		status = http.StatusMethodNotAllowed
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
