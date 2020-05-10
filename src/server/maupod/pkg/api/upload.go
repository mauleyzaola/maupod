package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
)

func (a *ApiServer) AudioFileUpload(p TransactionExecutorParams) (status int, result interface{}, err error) {
	file, header, err := p.r.FormFile("file")
	if err != nil {
		status = http.StatusBadRequest
		return
	}

	filename := filepath.Join(os.TempDir(), helpers.NewUUID())
	dest, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	if _, err = io.Copy(dest, file); err != nil {
		status = http.StatusInternalServerError
		return
	}
	if err = dest.Close(); err != nil {
		status = http.StatusInternalServerError
		return
	}
	defer func() {
		var localErr error
		if localErr = os.RemoveAll(filename); localErr != nil {
			log.Println(localErr)
		}
	}()

	var info *media.MediaInfo
	var ID string
	if err = media.MediaInfoWithId(filename, func(mi *media.MediaInfo, id string) {
		info = mi
		ID = id
	}); err != nil {
		status = http.StatusInternalServerError
		return
	}

	// store in db
	mi := info.ToDomain()
	mi.ID = ID
	mi.FileExtension = strings.ToLower(filepath.Ext(header.Filename))
	mi.Location = mi.ID + mi.FileExtension
	if err = a.dm.Insert(p.ctx, p.conn, mi); err != nil {
		status = http.StatusInternalServerError
		return
	}

	// TODO: upload to store

	result = mi
	status = http.StatusCreated

	return
}
