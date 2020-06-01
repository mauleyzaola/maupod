package main

import (
	"context"
	"errors"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/mauleyzaola/maupod/src/server/pkg/broker"
	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/filemgmt"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
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
	if allMedia, err = store.List(ctx, conn, data.MediaFilter{}, nil); err != nil {
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

	var timeout = time.Second * time.Duration(config.Delay)
	var output *pb.MediaInfoOutput

	for _, f := range files {
		input := &pb.MediaInfoInput{FileName: f}
		if output, err = broker.MediaInfoRequest(nc, input, timeout); err != nil {
			// TODO: send the files with errors to another listener and store in db
			logger.Error(err)
			continue
		}
		if output.Response == nil {
			return errors.New("missing response")
		}
		if !output.Response.Ok {
			return errors.New(output.Response.Error)
		}

		m := output.Media
		m.Id = helpers.NewUUID()
		// TODO: m.Sha  needs to be defined in another process
		m.LastScan = helpers.TimeToTs(&scanDate)
		m.ModifiedDate = output.LastModifiedDate
		m.Location = f
		m.FileExtension = filepath.Ext(f)

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

		// send message for extracting artwork if needed
		if rule.NeedsImageUpdate(m) {
			var payload []byte
			if payload, err = proto.Marshal(&pb.ArtworkExtractInput{Media: m, ScanDate: helpers.TimeToTs2(scanDate)}); err != nil {
				return err
			}
			if err = nc.Publish(strconv.Itoa(int(pb.Message_MESSAGE_ARTWORK_SCAN)), payload); err != nil {
				logger.Error(err)
			}
		}
	}

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
