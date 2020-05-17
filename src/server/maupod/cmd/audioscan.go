package cmd

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/filemgmt"

	"github.com/mauleyzaola/maupod/src/server/pkg/data/psql"
	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/spf13/cobra"
)

// scannerCmd represents the restapi command
var scannerCmd = &cobra.Command{
	Use:   "audioscan",
	Short: "Scans for new audio files",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := domain.ParseConfiguration()
		if err != nil {
			return err
		}

		if err = config.Validate(); err != nil {
			return err
		}

		var db *sql.DB
		//var conn *sql.Tx
		if db, err = helpers.DbBootstrap(config); err != nil {
			return err
		}
		defer func() {
			if err = db.Close(); err != nil {
				log.Println(err)
			}
		}()
		//defer func() {
		//	if conn != nil {
		//		if err != nil {
		//			err = conn.Rollback()
		//		} else {
		//			err = conn.Commit()
		//		}
		//		if err != nil {
		//			log.Println(err)
		//		}
		//	}
		//	if err = db.Close(); err != nil {
		//		log.Println(err)
		//	}
		//}()
		//if conn, err = db.Begin(); err != nil {
		//	return err
		//}

		// buffer all the existing hashes in db
		store := &psql.MediaStore{}
		ctx := context.Background()
		conn := db

		insertFn := func(ctx context.Context, ID, filename string, info *media.MediaInfo) error {
			// check if the same file exists
			ok, err := store.Exists(ctx, conn, ID)
			if err != nil {
				return err
			}
			if ok {
				// file already is stored in db
				return nil
			}
			m := info.ToDomain()
			m.FileExtension = filepath.Ext(filename)
			m.ID = ID
			m.Location = m.ID + m.FileExtension
			return store.Insert(ctx, conn, m)
		}

		var root string
		root = "/media/mau/music-library/music"
		//root = "/Volumes/Backup-Music-Library/music"
		//root = "/media/mau/music-library/music/Rush"
		//root = "/media/mau/music-library/music/Eric Clapton"
		//root = "/media/mau/music-library/music/Black Sabbath"
		// TODO: use a channel process to speed this up, reading files and sha takes too long

		if err = ScanFiles(ctx, root, config, insertFn); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scannerCmd)
}

type MediaInfoID struct {
	ID        string
	MediaInfo *media.MediaInfo
	Filename  string
}

func ScanFiles(ctx context.Context, root string, config *domain.Configuration,
	insertFn func(ctx context.Context, ID, filename string, info *media.MediaInfo) error) error {
	var err error
	var files []string
	const concurrentProcess = 10
	start := time.Now()

	walker := func(filename string, isDir bool) bool {
		if isDir {
			return false
		}
		if !config.FileIsValidExtension(filename) {
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

	tasks := make(chan string, concurrentProcess)
	results := make(chan MediaInfoID, concurrentProcess)

	var wg sync.WaitGroup
	wg.Add(len(files))
	go func() {
		for r := range results {
			if err := insertFn(ctx, r.ID, r.Filename, r.MediaInfo); err != nil {
				log.Println("[ERROR] ", err)
			}
			log.Printf("[SUCCESS] %s\n", filepath.Base(r.Filename))
			wg.Done()
		}
	}()

	for i, f := range files {
		if i%concurrentProcess == 0 {
			go ReadMediaInfoAsync(tasks, results)
		}
		tasks <- f
	}
	close(tasks)
	wg.Wait()

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}

// files only reads and medias only writes
func ReadMediaInfoAsync(files <-chan string, medias chan<- MediaInfoID) {
	for f := range files {
		log.Printf("[DEBUG] received file: %s\n", filepath.Base(f))
		_ = media.MediaInfoWithId(f, func(mi *media.MediaInfo, id string) {
			medias <- MediaInfoID{
				ID:        id,
				MediaInfo: mi,
				Filename:  f,
			}
		})
	}
}
