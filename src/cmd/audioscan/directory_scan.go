package main

import (
	"context"
	"log"
	"path/filepath"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/protos"
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
		cols.Sha,
	}
	return fields
}

func ScanDirectoryAudioFiles(
	ctx context.Context,
	conn boil.ContextExecutor,
	nc *nats.Conn,
	scanDate time.Time,
	store *dbdata.MediaStore,
	root string,
	config *protos.Configuration,
	force bool,
) error {

	var err error
	var files []string
	start := time.Now()

	// buffer all the media in db
	var allMedia dbdata.Medias
	if allMedia, err = store.List(ctx, conn, dbdata.MediaFilter{}, nil); err != nil {
		log.Println(err)
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

		if !force {
			var relativePath = paths.LocationPath(filename)
			// a bit of speed improvement, avoid a second time scanning the same file unless it has been changed in the file system
			if val, ok := mediaLocationKeys[relativePath]; ok {
				if !rules.NeedsMediaUpdate(val) {
					log.Println("[INFO] no need to scan this file: ", relativePath)
					return false
				}
			} else {
				log.Println("[INFO] location not found in db: ", relativePath)
			}
		}

		files = append(files, filename)
		return false
	}

	log.Println("[DEBUG] started scanning")
	if err = helpers.WalkFiles(paths.MediaFullPathAudioFile(root), walker); err != nil {
		log.Println(err)
		return err
	}

	var timeout = time.Second * time.Duration(config.Delay)
	for _, f := range files {
		var location = paths.LocationPath(f)
		var fullPath = paths.MediaFullPathAudioFile(location)
		var m *protos.Media
		var output *protos.MediaInfoOutput
		if output, err = broker.RequestMediaInfoScan(nc, fullPath, timeout); err != nil {
			log.Println(err)
			continue
		}
		m = output.Media

		// if the minimal information such as track, album and performer are not present, ignore this scan
		// TODO: send another NATS message for storing these errors
		if err = rules.MediaCheckMinimalData(m); err != nil {
			log.Println(err)
			continue
		}

		m.Id = helpers.NewUUID()
		m.LastScan = helpers.TimeToTs(&scanDate)
		m.Location = location
		m.Directory = filepath.Dir(m.Location)
		m.FileExtension = filepath.Ext(location)

		// if the location is the same and we made it here, that means we need to update the row
		if val, ok := mediaLocationKeys[location]; ok {
			m.Id = val.Id
			m.AlbumIdentifier = val.AlbumIdentifier
			if err = store.Update(ctx, conn, m, updatableFields()...); err != nil {
				return err
			}

		} else {
			// consider assigning album identifier (experimental feature only on new media)
			var albumIdentifier string
			if albumIdentifier, err = AlbumGroupDetection(ctx, conn, m); err != nil {
				log.Println(err)
				return err
			}
			m.AlbumIdentifier = albumIdentifier
			if err = store.Insert(ctx, conn, m); err != nil {
				log.Println(err)
				return err
			}
		}

		// send SHA update message
		if err = broker.PublishMediaSHAScan(nc, &protos.SHAScanInput{Media: m}); err != nil {
			log.Println(err)
			return err
		}

		// send message for extracting artwork if album has identifier
		// this will only look for image files in the same directory of the audio files
		// no scanning of audio files content should be done
		if m.AlbumIdentifier != "" {
			if err = broker.PublishBroker(nc, protos.Message_MESSAGE_ARTWORK_SCAN,
				&protos.ArtworkExtractInput{
					Media:    m,
					ScanDate: helpers.TimeToTs2(scanDate),
				}); err != nil {
				log.Println(err)
			}
		}
	}

	log.Printf("[INFO] files: %d  elapsed: %s\n", len(files), time.Since(start))

	return nil
}
