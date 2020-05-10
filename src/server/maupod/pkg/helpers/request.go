package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"

	"github.com/gorilla/mux"
)

func QueryStringValue(r *http.Request, name string) string {
	value := r.URL.Query().Get(name)
	if value != "" {
		return value
	}

	value = mux.Vars(r)[name]
	if value != "" {
		return value
	}
	return r.Header.Get(name)
}

func WriteJson(w http.ResponseWriter, err error, code int, v interface{}) {
	type DefaultApiMessage struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}

	// Set the content type.
	w.Header().Set("Content-Type", "application/json")

	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}

	if err != nil {
		if code == 0 {
			code = http.StatusInternalServerError
		}
		v = &DefaultApiMessage{
			Error: err.Error(),
			Ok:    false,
		}
	} else {
		if code == 0 {
			code = http.StatusOK
		}
		if v == nil {
			v = &DefaultApiMessage{
				Ok: true,
			}
		}
	}

	// Write the status code to the response and context.
	w.WriteHeader(code)

	// Marshal the data into a JSON string.
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("%s : Respond Marshalling JSON response\n", err)
	}
	if _, err = w.Write(data); err != nil {
		log.Println(err)
	}
}

func DecodeBody(r io.Reader, v interface{}) (fieldNames helpers.StringSlice, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, v); err != nil {
		return
	}

	var fieldNamesMap map[string]interface{}
	if err = json.Unmarshal(data, &fieldNamesMap); err != nil {
		return
	}

	for k, _ := range fieldNamesMap {
		fieldNames = append(fieldNames, k)
	}

	return fieldNames, nil
}
