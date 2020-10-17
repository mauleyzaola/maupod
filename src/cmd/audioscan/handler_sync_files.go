package main

import (
	"context"
	"fmt"
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

	// get the playedMediaList that needs to be sync
	store := dbdata.NewMediaStore()
	playedMediaList, err := store.PlayedMediaList(ctx, conn)
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

	// store the playedMediaList files we will process
	// the selection will depend on either choosing files or directories
	// so, we use a map to store the files and avoid double processing
	var sources, destinations []string
	var albumKey = make(map[string]struct{})

	for _, v := range playedMediaList {
		var medias []*pb.Media
		if input.IncludeDirectory {
			// ignore the file if the same album directory has been processed already
			if _, ok := albumKey[v.AlbumIdentifier]; ok {
				continue
			}
			// load the rest of the files from the same album
			// we are assuming same album files are in the same directory
			medias, err = store.FindMedias(ctx, conn, &pb.Media{AlbumIdentifier: v.AlbumIdentifier}, 0)
			if err != nil {
				log.Println(err)
				output.Error = err.Error()
				return
			}
		} else {
			medias = []*pb.Media{v}
		}

		for _, media := range medias {
			sources = append(sources, filepath.Join(srcDir, media.Location))
			destinations = append(destinations, filepath.Join(destDir, media.Location))
		}
		albumKey[v.AlbumIdentifier] = struct{}{}
	}

	if l1, l2 := len(sources), len(destinations); l1 != l2 {
		err = fmt.Errorf("sources: %d destinations: %d cannot process", l1, l2)
		log.Println(err)
		output.Error = err.Error()
		return
	}

	// execute sync for each playedMediaList file
	var fileCount int
	var now = time.Now()
	for i, src := range sources {
		dest := destinations[i]
		ok, localErr := syncFile(
			src,
			dest,
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
	log.Printf("[INFO] total played files: %d total included files: %d\n", len(playedMediaList), len(sources))
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
