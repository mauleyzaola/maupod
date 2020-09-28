package main

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerSyncFiles(msg *nats.Msg) {
	var input pb.SyncFilesInput
	var output pb.SyncFilesOutput
	var err error

	defer func() {
		if msg.Reply == "" {
			return
		}
		var data []byte
		if data, err = helpers.ProtoMarshal(&output); err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
		}
	}()

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		output.Error = err.Error()
		return
	}

	ctx := context.Background()
	conn := m.db

	// get the media that needs to be sync
	store := dbdata.NewMediaStore()
	media, err := store.PlayedMediaList(ctx, conn)
	if err != nil {
		output.Error = err.Error()
		return
	}

	// execute sync for each media file
	for _, v := range media {
		if err = syncFile(
			filepath.Join(paths.RootDirectory(), v.Location),
			filepath.Join(paths.SyncRootDirectory(), v.Location),
		); err != nil {
			log.Println(err)
			output.Error = err.Error()
			return
		}
	}
	return
}

// syncFile will copy one source file to destination if they are different
// on the modified date or the size
// or the destination file does not exist
func syncFile(src, dest string) error {
	var needsSync bool
	srcInfo, err := os.Stat(src)
	if err != nil {
		log.Println(err)
		return err
	}
	destInfo, err := os.Stat(dest)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println(err)
			return err
		}
		needsSync = true
	}
	if !needsSync {
		needsSync = srcInfo.Size() != destInfo.Size() ||
			srcInfo.ModTime().After(destInfo.ModTime())
	}
	if !needsSync {
		log.Printf("file: %s no sync needed\n", src)
		return nil
	}

	// make sure the dest directory exists
	dir := filepath.Dir(dest)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Println(err)
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if err = srcFile.Close(); err != nil {
			log.Println(err)
		}
	}()
	destFile, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if err = destFile.Close(); err != nil {
			log.Println(err)
		}
	}()
	byteCount, err := io.Copy(destFile, srcFile)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("[INFO] successfully copeid %d bytes to %s\n", byteCount, dest)

	return errors.New("not implemented")
}
