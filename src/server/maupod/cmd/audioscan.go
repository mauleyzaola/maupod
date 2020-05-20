package cmd

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/filemgmt"
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/spf13/cobra"
)

// scannerCmd represents the restapi command
var scannerCmd = &cobra.Command{
	Use:   "audioscan",
	Short: "Scans for new audio files",
	Long:  "Parameters should be directories where the audio files live. Edit config file to enable the file extensions as needed",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := rule.ConfigurationParse()
		if err != nil {
			return err
		}

		if err = rule.ConfigurationValidate(config); err != nil {
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
		conn := db

		var filter = filters.MediaFilter{}
		var allMedia data.Medias
		if allMedia, err = store.List(ctx, conn, filter, nil); err != nil {
			return err
		}

		mediaLocationKeys := allMedia.ToMap()
		var cols = orm.MediumColumns
		var fields = []string{
			cols.Sha,
			cols.FileExtension,
			cols.Format,
			cols.FileSize,
			cols.Duration,
			cols.OverallBitRateMode,
			cols.OverallBitRate,
			cols.StreamSize,
			cols.Album,
			cols.Track,
			cols.Title,
			cols.TrackPosition,
			cols.Performer,
			cols.Genre,
			cols.RecordedDate,
			cols.Comment,
			cols.Channels,
			cols.ChannelPositions,
			cols.ChannelLayout,
			cols.SamplingRate,
			cols.SamplingCount,
			cols.BitDepth,
			cols.CompressionMode,
			cols.EncodedLibrary,
			cols.EncodedLibraryName,
			cols.EncodedLibraryVersion,
			cols.BitRateMode,
			cols.BitRate,
			cols.LastScan,
			cols.ModifiedDate,
			cols.TrackNameTotal,
			cols.AlbumPerformer,
			cols.AudioCount,
			cols.BitDepthString,
			cols.CommercialName,
			cols.CompleteName,
			cols.CountOfAudioStreams,
			cols.EncodedLibraryDate,
			cols.FileName,
			cols.FolderName,
			cols.FormatInfo,
			cols.FormatURL,
			cols.InternetMediaType,
			cols.KindOfStream,
			cols.Part,
			cols.PartTotal,
			cols.StreamIdentifier,
			cols.WritingLibrary,
		}

		insertFn := func(ctx context.Context, filename string, info *media.MediaInfo) error {
			fileInfo, err := os.Stat(filename)
			if err != nil {
				return err
			}
			m := rule.NewMediaFile(info, filename, scanDate, fileInfo)

			// if the location is the same and we made it here, that means we need to update the row
			if val, ok := mediaLocationKeys[filename]; ok {
				m.Id = val.Id
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

func ScanFiles(ctx context.Context, root string, config *pb.Configuration,
	mediaLocationKeys map[string]*pb.Media,
	insertFn func(ctx context.Context, filename string, info *media.MediaInfo) error) error {
	var err error
	var files []string
	start := time.Now()

	walker := func(filename string, isDir bool) bool {
		if isDir {
			return false
		}
		if !rule.FileIsValidExtension(config, filename) {
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

			lastUpdateDb, lastUpdateFileSystem := val.ModifiedDate.Seconds, info.ModTime().Unix()
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
		info, err := media.MediaInfoFromFile(f)
		if err != nil {
			log.Printf("[ERROR] cannot get mediainfo from file: %s %s\n", f, err)
			continue
		}
		if err := insertFn(ctx, f, info); err != nil {
			log.Println("[ERROR] ", err)
		}
		//log.Printf("[SUCCESS] %s\n", filepath.Base(f))
	}

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
