package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/media"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"

	"github.com/nats-io/nats.go"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/filemgmt"
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/volatiletech/sqlboiler/boil"
)

func ScanDirectoryAudioFiles(
	ctx context.Context,
	conn boil.ContextExecutor,
	nc *nats.Conn,
	logger types.Logger,
	scanDate time.Time,
	store *data.MediaStore,
	root string,
	config *pb.Configuration,
) error {

	var err error
	var files []string
	start := time.Now()

	// buffer all the media in db
	var allMedia data.Medias
	if allMedia, err = store.List(ctx, conn, filters.MediaFilter{}, nil); err != nil {
		logger.Error(err)
		return err
	}

	mediaLocationKeys := allMedia.ToMap()

	walker := func(filename string, isDir bool) bool {
		if isDir {
			return false
		}
		if !rule.FileIsValidExtension(config, filename) {
			return false
		}

		// a bit of speed improvement, avoid a second time scanning the same file unless it has been changed in the file system
		if val, ok := mediaLocationKeys[filename]; ok {
			if !rule.NeedsUpdate(val) {
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

	var cols = orm.MediumColumns
	var fields = []string{
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
		cols.Composer,
	}

	for _, f := range files {
		mi, err := media.RunMediaInfo(f)
		if err != nil {
			// TODO: send the files with errors to another listener and store in db
			logger.Info("error using mediainfo with file: " + f)
			return err
			//continue
		}

		m := mi.ToProto()
		fileInfo, err := os.Stat(f)
		if err != nil {
			return err
		}
		m.Id = helpers.NewUUID()
		// TODO: m.Sha  needs to be defined in another process
		m.LastScan = helpers.TimeToTs(&scanDate)
		m.ModifiedDate = helpers.TimeToTs2(fileInfo.ModTime())
		m.Location = f
		m.FileExtension = filepath.Ext(fileInfo.Name())

		// if the location is the same and we made it here, that means we need to update the row
		if val, ok := mediaLocationKeys[f]; ok {
			m.Id = val.Id
			if err = store.Update(ctx, conn, m, fields...); err != nil {
				return err
			}
		} else {
			if err = store.Insert(ctx, conn, m); err != nil {
				return err
			}
		}
	}

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
