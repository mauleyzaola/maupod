package discogs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

// this value is set using a discogs account: https://www.discogs.com/settings/developers
// then, generate a new token
const consumerTokenEnvVarName = "CONSUMER_TOKEN"

const apiBaseURL = "https://api.discogs.com/database"

func tokenFromEnv() (*string, error) {
	token := viper.GetString(consumerTokenEnvVarName)
	if token == "" {
		return nil, errors.New("cannot resolve variable: " + consumerTokenEnvVarName)
	}
	return &token, nil
}

// Search documentation: https://www.discogs.com/developers#page:database,header:database-search
func Search(params url.Values) (*SearchResult, error) {
	token, err := tokenFromEnv()
	if err != nil {
		return nil, err
	}

	// TODO: set timeout
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	uri, err := url.Parse(apiBaseURL + "/search")
	if err != nil {
		return nil, err
	}

	// add the authentication to the request
	params.Set("token", *token)
	uri.RawQuery = params.Encode()

	request, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if expected, actual := http.StatusOK, response.StatusCode; expected != actual {
		return nil, fmt.Errorf("discogs api: wrong status code, expected: %v got: %v", expected, actual)
	}

	if response.Body == nil {
		return nil, errors.New("discogs api: missing body in response")
	}
	defer func() { _ = response.Body.Close() }()

	var result SearchResult
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

type SearchResult struct {
	Pagination Pagination `json:"pagination"`
	Results    []*Result  `json:"results"`
}

type Pagination struct {
	Page    int  `json:"page"`
	Pages   int  `json:"pages"`
	PerPage int  `json:"per_page"`
	Items   int  `json:"items"`
	Urls    Urls `json:"urls"`
}

type Urls struct {
	Last string `json:"last"`
	Next string `json:"next"`
}

type Result struct {
	// fake id just to make each result unique
	UUID string `json:"uuid"`

	Country        string    `json:"country"`
	Year           string    `json:"year"`
	Format         []string  `json:"format"`
	Label          []string  `json:"label"`
	Type           string    `json:"type"`
	Genre          []string  `json:"genre"`
	Style          []string  `json:"style"`
	ID             int       `json:"id"`
	Barcode        []string  `json:"barcode"`
	UserData       UserData  `json:"user_data"`
	MasterID       int       `json:"master_id"`
	MasterURL      string    `json:"master_url"`
	URI            string    `json:"uri"`
	CatNo          string    `json:"catno"`
	Title          string    `json:"title"`
	Thumb          string    `json:"thumb"`
	CoverImage     string    `json:"cover_image"`
	ResourceURL    string    `json:"resource_url"`
	Community      Community `json:"community"`
	FormatQuantity int       `json:"format_quantity,omitempty"`
	Formats        []Format  `json:"formats,omitempty"`
}

type UserData struct {
	InWantlist   bool `json:"in_wantlist"`
	InCollection bool `json:"in_collection"`
}

type Community struct {
	Want int `json:"want"`
	Have int `json:"have"`
}

type Format struct {
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
	Descriptions []string `json:"descriptions"`
}
