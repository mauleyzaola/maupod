package cmd

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/psql"
	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/mauleyzaola/maupod/src/server/pkg/filemgmt"
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/spf13/cobra"
)

// scannerCmd represents the restapi command
var scannerCmd = &cobra.Command{
	Use:   "audioscan",
	Short: "Scans for new audio files",
	Long:  "Parameters should be directories where the audio files live. Edit config file to enable the file extensions as needed",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := domain.ParseConfiguration()
		if err != nil {
			return err
		}

		if err = config.Validate(); err != nil {
			return err
		}

		if len(os.Args) < 3 {
			return errors.New("missing path for file scan")
		}

		roots := os.Args[2:]
		for _, root := range roots {
			if _, err = os.Stat(root); err != nil {
				return err
			}
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
		var allMedia domain.Medias
		if allMedia, err = store.List(ctx, conn, filter, nil); err != nil {
			return err
		}

		mediaLocationKeys := allMedia.ToMap()
		var cols = orm.MediumColumns
		var fields = []string{cols.ModifiedDate, cols.LastScan, cols.Sha, cols.FileExtension, cols.Duration,
			cols.BitRate, cols.BitRateMode, cols.EncodedLibraryVersion, cols.EncodedLibrary, cols.EncodedLibraryName,
			cols.Format, cols.FileSize, cols.OverallBitRateMode, cols.OverallBitRate, cols.StreamSize, cols.Album, cols.Track,
			cols.Title, cols.TrackPosition, cols.Performer, cols.Genre, cols.RecordedDate, cols.FileModifiedDate, cols.Comment,
			cols.Channels, cols.ChannelPositions, cols.ChannelLayout, cols.SamplingRate, cols.SamplingCount, cols.BitDepth,
			cols.CompressionMode}

		insertFn := func(ctx context.Context, filename string, info *media.MediaInfo) error {
			fileInfo, err := os.Stat(filename)
			if err != nil {
				return err
			}
			m := rule.NewMediaFile(info, filename, scanDate, fileInfo)

			// if the location is the same and we made it here, that means we need to update the row
			if val, ok := mediaLocationKeys[filename]; ok {
				m.ID = val.ID
				return store.Update(ctx, conn, m, fields)
			} else {
				return store.Insert(ctx, conn, m)
			}
		}

		for _, root := range roots {
			if err = ScanFiles(ctx, root, config, mediaLocationKeys, insertFn); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scannerCmd)
}

func ScanFiles(ctx context.Context, root string, config *domain.Configuration,
	mediaLocationKeys map[string]*domain.Media,
	insertFn func(ctx context.Context, filename string, info *media.MediaInfo) error) error {
	var err error
	var files []string
	start := time.Now()

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

			lastUpdateDb, lastUpdateFileSystem := val.ModifiedDate.Unix(), info.ModTime().Unix()
			diffSeconds := math.Abs(float64(lastUpdateFileSystem - lastUpdateDb))

			// less than 5 seconds should not be considered as a change
			if diffSeconds < 5 {
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

	for _, f := range files {
		infos, err := media.MediaInfoFromFiles(f)
		if err != nil {
			log.Printf("[ERROR] cannot get mediainfo from file: %s %s\n", f, err)
			continue
		}
		if len(infos) != 1 {
			log.Println("[ERROR] infos is more than one:", f)
			continue
		}
		info := &infos[0]

		if err := insertFn(ctx, f, info); err != nil {
			log.Println("[ERROR] ", err)
		}
		//log.Printf("[SUCCESS] %s\n", filepath.Base(f))
	}

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
