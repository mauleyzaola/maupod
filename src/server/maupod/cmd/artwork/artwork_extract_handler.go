package main

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/images"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

const thumbnailDir = "thumbnail"

func (m *MsgHandler) handlerArtworkExtract(msg *nats.Msg) {
	var input pb.ArtworkExtractInput
	err := proto.Unmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}

	imageData, err := ScanArtwork(m.base.Logger(), input.Media)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}

	// we set the last image scan in any case, so we don't pass a second time this file, unless there are changes
	input.Media.LastImageScan = input.ScanDate

	// save the image file to file store
	if imageData != nil && input.Media.ShaImage != "" {
		ctx := context.Background()
		conn := m.db
		store := &data.MediaStore{}

		var matchedMedias []*pb.Media
		var matchedMedia *pb.Media

		if matchedMedias, err = store.FindMedias(ctx, conn, &pb.Media{
			ShaImage: input.Media.ShaImage,
		}, 2); err != nil {
			m.base.Logger().Error(err)
			return
		}
		for _, v := range matchedMedias {
			if v.Id != input.Media.Id {
				matchedMedia = v
				input.Media.ImageLocation = v.ImageLocation
				m.base.Logger().Infof("found a matched sha image for audio file: %s", input.Media.Id)
				break
			}
		}

		// if there was an image in the audio file...
		if matchedMedia == nil {
			// write artwork if there is no other media with the same sha image
			var x, y int
			input.Media.ImageLocation = rule.ArtworkFileName(input.Media)
			if x, y, err = images.Size(bytes.NewBuffer(imageData)); err != nil {
				m.base.Logger().Error(err)
				return
			}
			// check only square images are allowed
			if x == y {
				// try to generate the big file
				if x >= int(m.config.ArtworkBigSize) {
					if err = images.ImageResize(
						bytes.NewBuffer(imageData),
						filepath.Join(m.config.ArtworkStore.Location, input.Media.ImageLocation),
						int(m.config.ArtworkBigSize),
						int(m.config.ArtworkBigSize),
					); err != nil {
						m.base.Logger().Error(err)
						return
					}
				}
				// try to generate the small file
				if x >= int(m.config.ArtworkSmallSize) {
					if err = os.MkdirAll(filepath.Join(m.config.ArtworkStore.Location, thumbnailDir), os.ModePerm); err != nil {
						m.base.Logger().Error(err)
						return
					}
					if err = images.ImageResize(
						bytes.NewBuffer(imageData),
						filepath.Join(m.config.ArtworkStore.Location, thumbnailDir, input.Media.ImageLocation),
						int(m.config.ArtworkSmallSize),
						int(m.config.ArtworkSmallSize),
					); err != nil {
						m.base.Logger().Error(err)
						return
					}
				}
			}
		}
	}

	payload, err := proto.Marshal(&pb.ArtworkUpdateInput{Media: input.Media})
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	if err = m.base.NATS().Publish(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_UPDATE_ARTWORK)), payload); err != nil {
		m.base.Logger().Error(err)
		return
	}
}

// ScanArtwork will take one audio media file as parameter and:
// 1. try to extract the image from the audio file
// 2. get the sha of the image
// in any case, the media last_scan_date will be updated
// an error will return if something critical fails in the process
func ScanArtwork(
	logger types.Logger,
	media *pb.Media,
) ([]byte, error) {
	// check media last_image_scan vs date last modified
	if !rule.NeedsImageUpdate(media) {
		logger.Info("image does not need to be updated")
		return nil, nil
	}

	// we should not process the image a second time here, this is just a file scan
	if rule.MediaHasImage(media) {
		logger.Info("image does not have image data")
		return nil, nil
	}

	var err error
	var shaData []byte
	var imageData []byte

	w := &bytes.Buffer{}
	if err = images.ExtractImageFromMedia(w, media.Location); err != nil {
		// no image in audio file, update scan in db and return
		logger.Info("no image found in audio file: " + media.Location)
		return nil, nil
	} else {
		imageData = w.Bytes()
	}

	if imageData != nil {
		// calculate the sha of the image
		if shaData, err = helpers.SHA(bytes.NewBuffer(imageData)); err != nil {
			logger.Error(err)
			return nil, err
		}
		media.ShaImage = helpers.HashFromSHA(shaData)
	}

	return imageData, nil
}
