package cmd

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/filters"

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

		root := os.Args[len(os.Args)-1]
		if _, err = os.Stat(root); err != nil {
			return err
		}
		scanDate := time.Now()

		var db *sql.DB
		if db, err = helpers.DbBootstrap(config); err != nil {
			return err
		}
		defer func() {
			if err = db.Close(); err != nil {
				log.Println(err)
			}
		}()

		store := &psql.MediaStore{}
		ctx := context.Background()
		conn := db

		var filter = filters.MediaFilter{}
		var allMedia []*domain.Media
		if allMedia, err = store.List(ctx, conn, filter, nil); err != nil {
			return err
		}

		insertFn := func(ctx context.Context, ID, filename string, info *media.MediaInfo, fileInfo os.FileInfo) error {
			// check if the same file exists based on the ID = sha256
			ok, err := store.Exists(ctx, conn, ID)
			if err != nil {
				return err
			}
			if ok {
				// file already is stored in db
				return nil
			}

			m := info.ToDomain()
			m.ID = ID
			m.FileExtension = filepath.Ext(filename)
			m.Location = filename
			m.LastScan = scanDate
			m.ModifiedDate = fileInfo.ModTime()

			return store.Insert(ctx, conn, m)
		}

		if err = ScanFiles(ctx, root, config, allMedia, insertFn); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scannerCmd)
}

// MediaInfoID consolidates the data that flows through the channels
type MediaInfoID struct {
	ID        string
	Filename  string
	MediaInfo *media.MediaInfo
	FileInfo  os.FileInfo
}

func ScanFiles(ctx context.Context, root string, config *domain.Configuration,
	allMedia domain.Medias,
	insertFn func(ctx context.Context, ID, filename string, info *media.MediaInfo, fileInfo os.FileInfo) error) error {
	var err error
	var files []string
	const concurrentProcess = 10
	start := time.Now()

	mediaLocationKeys := allMedia.ToMap()

	walker := func(filename string, isDir bool) bool {
		if isDir {
			return false
		}
		if !config.FileIsValidExtension(filename) {
			return false
		}

		// a bit of speed improvement, avoid a second time scanning the same file unless it has been changed in the file system
		if val, ok := mediaLocationKeys[filename]; ok {
			// TODO: avoid a second file scan?
			info, err := os.Stat(filename)
			if err != nil {
				log.Println("[ERROR] ", err)
				return false
			}
			if val.ModifiedDate.Before(info.ModTime()) {
				return false
			}
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
			if err := insertFn(ctx, r.ID, r.Filename, r.MediaInfo, r.FileInfo); err != nil {
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
		info, err := os.Stat(f)
		if err != nil {
			log.Println("[ERROR] ", err)
			continue
		}

		_ = media.MediaInfoWithId(f, func(mi *media.MediaInfo, id string) {
			medias <- MediaInfoID{
				ID:        id,
				MediaInfo: mi,
				Filename:  f,
				FileInfo:  info,
			}
		})
	}
}
