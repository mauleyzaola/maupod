package api

import (
	"fmt"
	"io"
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

	//log.Printf("received file: %s", header.Filename)
	//log.Printf("file size: %v", header.Size)

	id := helpers.NewUUID()
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
		if err = dest.Close(); err != nil {
			log.Println(err)
		}
	}()

	if _, err = io.Copy(dest, file); err != nil {
		status = http.StatusInternalServerError
		return
	}
	//log.Println("wrote destination file to: ", filename)

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

	// TODO: check for duplicated files, based on the mediainfo they contain, not on the path to the file
	// TODO: set location to be something else
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
