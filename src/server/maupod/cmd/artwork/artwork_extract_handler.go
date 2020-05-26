package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/types"

	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/images"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/boil"
)

func (m *MsgHandler) handlerArtworkExtract(msg *nats.Msg) {
	var input pb.ArtworkExtractInput
	err := proto.Unmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	m.base.Logger().Info("received artwork extract message: " + input.String())

	ctx := context.Background()
	conn := m.db

	if err = ScanArtwork(ctx, conn, m.base.Logger(), &data.MediaStore{}, helpers.TsToTime2(input.ScanDate), input.Media, m.imageStore); err != nil {
		m.base.Logger().Error(err)
		return
	}
}

// ScanArtwork will take one audio media file as parameter and:
// 1. try to extract the image from the audio file
// 2. get the sha of the image and compare to existing sha in another files
// 3. if the sha is unique, save the image data to store
// 4. if not unique, will copy the sha image and the location from existing media
// in any case, the media last_scan_date will be updated
func ScanArtwork(
	ctx context.Context,
	conn boil.ContextExecutor,
	logger types.Logger,
	store *data.MediaStore,
	scanDate time.Time,
	media *pb.Media,
	imageStore *pb.FileStore,
) error {

	var err error

	// check media last_image_scan vs date last modified
	if !rule.NeedsImageUpdate(media) {
		return nil
	}

	var shaData []byte
	var imageData []byte
	var saveImage bool
	var filename string
	var cols = orm.MediumColumns

	// extract the image from the file
	media.LastImageScan = helpers.TimeToTs(&scanDate)
	w := &bytes.Buffer{}
	if err = images.ExtractImageFromMedia(w, media.Location); err != nil {
		// no image in audio file, update scan in db and continue
		if err = store.Update(ctx, conn, media, cols.LastImageScan); err != nil {
			logger.Error(err)
		}
		return err
	} else {
		imageData = w.Bytes()
	}

	if imageData != nil {
		// calculate the sha of the image
		if shaData, err = helpers.SHA(bytes.NewBuffer(imageData)); err != nil {
			logger.Error(err)
			return err
		}
		media.ShaImage = helpers.HashFromSHA(shaData)

		// see if there is another media with the same shaimage
		var medias []*pb.Media
		if medias, err = store.FindMedias(ctx, conn, &pb.Media{ShaImage: media.ShaImage}, 2); err != nil {
			logger.Error(err)
			return err
		}
		var mediaSameShaImage *pb.Media
		for _, v := range medias {
			if v.Id != media.Id {
				mediaSameShaImage = v
				break
			}
		}
		if mediaSameShaImage != nil {
			media.ImageLocation = mediaSameShaImage.ImageLocation
			saveImage = false
		} else {
			saveImage = true
		}
	} else {
		media.ImageLocation = ""
		media.ShaImage = ""
	}

	if saveImage {
		if filename, err = rule.ImageFileName(media, imageStore); err != nil {
			logger.Error(err)
			return err
		}
		media.ImageLocation = filename
		// TODO: resize image, for now store the original
		if err = ioutil.WriteFile(filename, imageData, os.ModePerm); err != nil {
			logger.Error(err)
			return err
		}
	}

	// update media record
	if err = store.Update(ctx, conn, media,
		cols.LastImageScan,
		cols.ImageLocation,
		cols.ShaImage,
	); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
