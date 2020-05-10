package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
)

func (a *ApiServer) AudioFileUpload(p TransactionExecutorParams) (status int, result interface{}, err error) {
	file, header, err := p.r.FormFile("file")
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	// generate the hash from the file itself, which will be the id in db and path
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	buffer := bytes.NewBuffer(fileData)
	hash, err := helpers.SHA(buffer)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	id := fmt.Sprintf("%x", string(hash))

	log.Println("hash: ", id)

	// TODO: move to the real location in s3
	targetDir := os.TempDir()
	ext := filepath.Ext(header.Filename)
	filename := id + ext
	filename = filepath.Join(targetDir, filename)
	dest, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	defer func() {
		var localErr error
		if localErr = dest.Close(); localErr != nil {
			log.Println(localErr)
		}
		if localErr = os.RemoveAll(filename); localErr != nil {
			log.Println(localErr)
		}
	}()

	buffer = bytes.NewBuffer(fileData)
	if _, err = io.Copy(dest, buffer); err != nil {
		status = http.StatusInternalServerError
		return
	}

	// extract mediainfo from destination
	infos, err := media.MediaInfoFromFiles(filename)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	if len(infos) != 1 {
		err = fmt.Errorf("expected media files to be: 1 instead got: %v", len(infos))
		status = http.StatusBadRequest
		return
	}
	info := infos[0]

	// store in db
	mi := info.ToDomain()
	mi.ID = id
	mi.Location = filepath.Base(filename)
	if err = a.dm.Insert(p.ctx, p.conn, mi); err != nil {
		status = http.StatusInternalServerError
		return
	}

	// TODO: upload to store

	result = mi
	status = http.StatusCreated

	return
}
