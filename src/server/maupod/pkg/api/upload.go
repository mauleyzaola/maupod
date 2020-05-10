package api

import (
	"io"
	"log"
	"net/http"
	"os"
)

func (a *ApiServer) AudioFileUpload(p TransactionExecutorParams) (status int, result interface{}, err error) {
	file, header, err := p.r.FormFile("file")
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	log.Printf("received file: %s", header.Filename)
	log.Printf("file size: %v", header.Size)

	dest, err := os.OpenFile("/home/mau/Downloads/destination", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	defer func() {
		if err = dest.Close(); err != nil {
			log.Println(err)
		}
	}()

	if _, err = io.Copy(dest, file); err != nil {
		status = http.StatusInternalServerError
		return
	}

	result = struct {
		Beer interface{} `json:"beer"`
	}{
		"is good",
	}
	return
}
