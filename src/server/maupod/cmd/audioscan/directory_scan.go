package main

import (
	"context"
	"path/filepath"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/broker"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rules"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/boil"
)

func updatableFields() []string {
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
	return fields
}

func ScanDirectoryAudioFiles(
	ctx context.Context,
	conn boil.ContextExecutor,
	nc *nats.Conn,
	logger types.Logger,
	scanDate time.Time,
	store *dbdata.MediaStore,
	root string,
	config *pb.Configuration,
) error {

	var err error
	var files []string
	start := time.Now()

	// buffer all the media in db
	var allMedia dbdata.Medias
	if allMedia, err = store.List(ctx, conn, dbdata.MediaFilter{}, nil); err != nil {
		logger.Error(err)
		return err
	}

	mediaLocationKeys := allMedia.ToMap()

	walker := func(filename string, isDir bool) bool {
		if isDir {
			return false
		}
		if !rules.FileIsValidExtension(config, filename) {
			return false
		}

		// a bit of speed improvement, avoid a second time scanning the same file unless it has been changed in the file system
		if val, ok := mediaLocationKeys[filename]; ok {
			if !rules.NeedsMediaUpdate(val) {
				return false
			}
		}

		files = append(files, filename)
		return false
	}

	logger.Info("[DEBUG] started scanning")
	if err = helpers.WalkFiles(root, walker); err != nil {
		logger.Error(err)
		return err
	}

	var timeout = time.Second * time.Duration(config.Delay)
	for _, f := range files {
		var m *pb.Media
		if m, err = broker.RequestMediaInfoScan(nc, logger, f, timeout); err != nil {
			logger.Error(err)
			continue
		}

		m.Id = helpers.NewUUID()
		m.LastScan = helpers.TimeToTs(&scanDate)
		m.Location = f
		m.FileExtension = filepath.Ext(f)

		// if the location is the same and we made it here, that means we need to update the row
		if val, ok := mediaLocationKeys[f]; ok {
			m.Id = val.Id
			if err = store.Update(ctx, conn, m, updatableFields()...); err != nil {
				return err
			}
		} else {
			// consider assigning album identifier (experimental feature only on new media)
			var albumIdentifier string
			var isCompilation bool
			if albumIdentifier, isCompilation, err = AlbumGroupDetection(ctx, conn, m); err != nil {
				return err
			}
			m.AlbumIdentifier = albumIdentifier
			m.IsCompilation = isCompilation
			if err = store.Insert(ctx, conn, m); err != nil {
				return err
			}
		}

		// send message for extracting artwork if needed
		if rules.NeedsImageUpdate(m) {
			var payload []byte
			if payload, err = helpers.ProtoMarshal(&pb.ArtworkExtractInput{Media: m, ScanDate: helpers.TimeToTs2(scanDate)}); err != nil {
				return err
			}
			if err = broker.PublishMessage(nc, pb.Message_MESSAGE_ARTWORK_SCAN, payload); err != nil {
				logger.Error(err)
			}
		}

		// on each case media has probably changed and we need to get the new sha hash
		if err = broker.PublishMediaSHAUpdate(nc, &pb.MediaInfoInput{Media: m, FileName: f}); err != nil {
			logger.Error(err)
			return err
		}
	}

	logger.Infof("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
