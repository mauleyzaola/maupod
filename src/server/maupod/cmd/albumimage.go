package cmd

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/filemgmt"
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/images"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/spf13/cobra"
)

var albumimageCmd = &cobra.Command{
	Use:   "albumimage",
	Short: "extract image from audio file",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := rule.ConfigurationParse()
		if err != nil {
			return err
		}

		if err = rule.ConfigurationValidate(config); err != nil {
			return err
		}

		// read first image store in the configuration and all the stores where to look up for audio files
		var imageStore = rule.ConfigurationFirstImageStore(config)
		if imageStore == nil {
			return errors.New("could not find any image store in configuration, exiting now")
		}

		var fileSystemStores = rule.ConfigurationFileSystemStores(config)
		var roots []string
		for _, v := range fileSystemStores {
			roots = append(roots, v.Location)
		}
		for _, root := range roots {
			if _, err = os.Stat(root); err != nil {
				return err
			}
		}

		scanDate := time.Now()
		var db *sql.DB
		if db, err = data.DbBootstrap(config); err != nil {
			return err
		}
		defer func() {
			if err = db.Close(); err != nil {
				log.Println(err)
			}
		}()

		store := &data.MediaStore{}
		ctx := context.Background()

		// for image handling, we need to consider db transactions to avoid wrong/incomplete data to be stored
		conn, err := db.Begin()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				err = conn.Rollback()
			} else {
				err = conn.Commit()
			}
			if err != nil {
				log.Fatal(err)
			}
		}()
		var filter = filters.MediaFilter{}
		var allMedia data.Medias
		if allMedia, err = store.List(ctx, conn, filter, nil); err != nil {
			return err
		}

		mediaLocationKeys := allMedia.ToMap()
		var cols = orm.MediumColumns
		var fields = []string{cols.LastImageScan, cols.ShaImage, cols.ImageLocation}
		updateFn := func(ctx context.Context, media *pb.Media) error {
			return store.Update(ctx, conn, media, fields)
		}

		for _, root := range roots {
			if err = ScanArtwork(ctx, scanDate, root, config, imageStore, mediaLocationKeys, updateFn); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(albumimageCmd)
}

func ScanArtwork(ctx context.Context,
	scanDate time.Time,
	root string,
	config *pb.Configuration,
	imageStore *pb.FileStore,
	mediaLocationKeys map[string]*pb.Media,
	updateFn func(ctx context.Context, info *pb.Media) error) error {

	var err error
	var files []string
	start := time.Now()

	// create directory if not exists
	if err = os.MkdirAll(imageStore.Location, os.ModePerm); err != nil {
		return err
	}

	// only consider files which don't have sha image yet
	shaImageKeys := make(map[string]struct{})

	// add which files we will extract the image from
	walker := func(filename string, isDir bool) bool {
		if isDir {
			return false
		}
		if !rule.FileIsValidExtension(config, filename) {
			return false
		}

		val, ok := mediaLocationKeys[filename]
		if !ok {
			// if the file has not yet been scanned, we cannot process its image
			return false
		}
		me := rule.Media(*val)

		if me.HasImage() {
			// if the file has already sha for the image, don't add it to the files
			// however we add it to the scanned images, so we don't extract the image a second time
			if _, ok = shaImageKeys[val.ShaImage]; !ok {
				shaImageKeys[val.ShaImage] = struct{}{}
			}
			return false
		}

		files = append(files, filename)
		return false
	}

	log.Println("[DEBUG] started scanning")
	if err = filemgmt.WalkFiles(root, walker); err != nil {
		log.Println(err)
		return err
	}
	log.Printf("[DEBUG] finished scanning %d files\n", len(files))

	// extract the images from each file
	for _, f := range files {
		var shaData []byte
		var imageData []byte
		var saveImage bool
		var filename string
		val, ok := mediaLocationKeys[f]
		if !ok {
			// this should not happen, as we have already skipped these cases in the walker
			continue
		}
		me := rule.Media(*val)
		if !me.NeedsImageUpdate() {
			continue
		}

		// extract the image from the file
		val.LastImageScan = helpers.TimeToTs(&scanDate)
		w := &bytes.Buffer{}
		if err = images.ExtractImageFromMedia(w, me.Location); err == nil {
			imageData = w.Bytes()
			log.Println("[DEBUG] found image data for file: ", me.FileName)

			// calculate the sha of the image
			if shaData, err = helpers.SHA(bytes.NewBuffer(imageData)); err != nil {
				// we should not allow this, simply ignore for now
			} else {
				val.ShaImage = helpers.HashFromSHA(shaData)
				log.Println("[DEBUG] hash ok for file length: ", len(val.ShaImage))
				if _, ok = shaImageKeys[val.ShaImage]; !ok {
					shaImageKeys[val.ShaImage] = struct{}{}
					// flag for later
					saveImage = true
				}
			}
		}

		// update database
		val.ImageLocation = filename
		if filename, err = rule.ImageFileName(val, imageStore); err != nil {
			return err
		}
		if err = updateFn(ctx, val); err != nil {
			return err
		}

		// if hash doesn't exist save to image store
		if saveImage {
			log.Println("[DEBUG] storing artwork at: ", filename)
			// TODO: resize image, for now store the original
			if err = ioutil.WriteFile(filename, imageData, os.ModePerm); err != nil {
				return err
			}
		}
	}

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
