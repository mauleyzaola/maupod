package discogs

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/spf13/viper"
)

func TestSearch(t *testing.T) {
	t.Skip("don't waste resources :D")
	viper.AutomaticEnv()

	_, err := tokenFromEnv()
	if err != nil {
		t.Skip(err)
	}

	var params = &url.Values{}
	params.Add("title", "music for the masses")
	params.Add("year", "1987")

	result, err := Search(*params)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	t.Logf("[INFO] found %d results", len(result.Results))

	for _, v := range result.Results {
		t.Logf("cover: %s", v.CoverImage)
	}
}
