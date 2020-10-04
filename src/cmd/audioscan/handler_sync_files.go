package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

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

	srcDir, destDir := paths.RootDirectory(), paths.SyncRootDirectory()

	if _, err = os.Stat(srcDir); err != nil {
		log.Println(err)
		output.Error = err.Error()
		return
	}

	if _, err = os.Stat(destDir); err != nil {
		log.Println(err)
		output.Error = err.Error()
		return
	}

	// execute sync for each media file
	var fileCount int
	var now = time.Now()
	for _, v := range media {
		ok, localErr := syncFile(
			filepath.Join(srcDir, v.Location),
			filepath.Join(destDir, v.Location),
		)
		if localErr != nil {
			log.Println(localErr)
			output.Error = localErr.Error()
			return
		}
		if ok {
			fileCount++
		}
	}
	log.Printf("[INFO] completed sync for: %d files elapsed: %s", fileCount, time.Since(now))
	return
}

// syncFile will copy one source file to destination if they are different
// on the modified date or the size
// or the destination file does not exist
func syncFile(src, dest string) (bool, error) {
	var needsSync bool
	srcInfo, err := os.Stat(src)
	if err != nil {
		log.Println(err)
		return false, err
	}
	destInfo, err := os.Stat(dest)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println(err)
			return false, err
		}
		needsSync = true
	}
	if !needsSync {
		needsSync = srcInfo.Size() != destInfo.Size() ||
			srcInfo.ModTime().After(destInfo.ModTime())
	}
	if !needsSync {
		//log.Printf("file: %s no sync needed\n", src)
		return false, err
	}

	// make sure the dest directory exists
	dir := filepath.Dir(dest)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Println(err)
		return false, err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer func() {
		if err = srcFile.Close(); err != nil {
			log.Println(err)
		}
	}()
	destFile, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer func() {
		if err = destFile.Close(); err != nil {
			log.Println(err)
		}
	}()
	byteCount, err := io.Copy(destFile, srcFile)
	if err != nil {
		log.Println(err)
		return false, err
	}
	log.Printf("[INFO] successfully copied %d bytes to %s\n", byteCount, dest)

	return true, nil
}
