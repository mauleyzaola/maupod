package main

import (
	"bytes"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/images"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerArtworkExtract(msg *nats.Msg) {
	var input pb.ArtworkExtractInput
	err := proto.Unmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	m.base.Logger().Info("received artwork extract message: " + input.String())

	if err = ScanArtwork(m.base.Logger(), helpers.TsToTime2(input.ScanDate), input.Media); err != nil {
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
	scanDate time.Time,
	media *pb.Media,
) error {
	// check media last_image_scan vs date last modified
	if !rule.NeedsImageUpdate(media) {
		return nil
	}

	// we should not process the image a second time here, this is just a file scan
	if rule.MediaHasImage(media) {
		return nil
	}

	var err error
	var shaData []byte
	var imageData []byte

	// extract the image from the file
	media.LastImageScan = helpers.TimeToTs(&scanDate)

	w := &bytes.Buffer{}
	if err = images.ExtractImageFromMedia(w, media.Location); err != nil {
		// no image in audio file, update scan in db and return
		return nil
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
	}

	return nil
}
