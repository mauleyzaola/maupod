package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/filemgmt"
	"github.com/mauleyzaola/maupod/src/server/pkg/filters"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/volatiletech/sqlboiler/boil"
)

func ScanDirectoryAudioFiles(
	ctx context.Context,
	conn boil.ContextExecutor,
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
		cols.Composer,
	}

	for _, f := range files {
		info, err := media.InfoFromFile(f)
		if err != nil {
			log.Printf("[ERROR] cannot get mediainfo from file: %s %s\n", f, err)
			continue
		}
		fileInfo, err := os.Stat(f)
		if err != nil {
			return err
		}
		m := rule.NewMediaFile(info, f, scanDate, fileInfo)

		// if the location is the same and we made it here, that means we need to update the row
		if val, ok := mediaLocationKeys[f]; ok {
			m.Id = val.Id
			return store.Update(ctx, conn, m, fields...)
		} else {
			return store.Insert(ctx, conn, m)
		}
	}

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
