package main

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata"
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

	var err error
	var input pb.ArtworkExtractInput
	if err = proto.Unmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		return
	}

	// obtain the image file from the media
	imageData, err := ScanArtwork(m.base.Logger(), input.Media)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}

	if imageData != nil {
		m.base.Logger().Infof("found image: %s", input.Media.Location)
	} else {
		m.base.Logger().Infof("image not found: %s", input.Media.Location)
	}

	// we need to update the database when this function exits, one way or another
	defer func() {
		var payload []byte
		if err != nil {
			m.base.Logger().Errorf("[ERROR]  %s", err.Error())
			return
		}
		if payload, err = proto.Marshal(&pb.ArtworkUpdateInput{Media: input.Media}); err != nil {
			m.base.Logger().Error(err)
			return
		}
		if err = m.base.NATS().Publish(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_UPDATE_ARTWORK)), payload); err != nil {
			m.base.Logger().Error(err)
			return
		}
	}()

	// we set the last image scan in any case, so we don't pass a second time this file, unless there are changes
	input.Media.LastImageScan = input.ScanDate

	// no image to play with, do nothing
	if imageData == nil || input.Media.ShaImage == "" {
		return
	}

	ctx := context.Background()
	conn := m.db
	store := &dbdata.MediaStore{}

	var matchedMedias []*pb.Media
	var matchedMedia *pb.Media

	// we only need to query two rows, so one has a different id than this media file
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
			break
		}
	}

	// if there was a match in the sha image with another file, exit
	if matchedMedia != nil {
		return
	}

	// write artwork if there is no other media with the same sha image
	var x, y int
	input.Media.ImageLocation = rule.ArtworkFileName(input.Media)
	if x, y, err = images.Size(bytes.NewBuffer(imageData)); err != nil {
		m.base.Logger().Error(err)
		input.Media.Location = ""
		input.Media.ShaImage = ""
		return
	}

	// check only square images are allowed
	if x != y {
		input.Media.Location = ""
		input.Media.ShaImage = ""
		m.base.Logger().Warningf("image has invalid shape: %vx%v\n", x, y)
		return
	}

	// if the image is not bigSize enough, exit
	if x < int(m.config.ArtworkBigSize) {
		input.Media.Location = ""
		input.Media.ShaImage = ""
		m.base.Logger().Warningf("image too small for artwork: %vx%v\n", x, y)
		return
	}

	// generate the bigSize file
	bigSize, smallSize := int(m.config.ArtworkBigSize), int(m.config.ArtworkSmallSize)
	if err = images.ImageResize(
		bytes.NewBuffer(imageData),
		filepath.Join(m.config.ArtworkStore.Location, input.Media.ImageLocation),
		bigSize,
		bigSize,
	); err != nil {
		m.base.Logger().Error(err)
		return
	}

	// generate the smallSize file
	if err = os.MkdirAll(filepath.Join(m.config.ArtworkStore.Location, thumbnailDir), os.ModePerm); err != nil {
		m.base.Logger().Error(err)
		return
	}
	if err = images.ImageResize(
		bytes.NewBuffer(imageData),
		filepath.Join(m.config.ArtworkStore.Location, thumbnailDir, input.Media.ImageLocation),
		smallSize,
		smallSize,
	); err != nil {
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

	// first attempt is picking the image from the same directory
	if imageData, _ = FindArtworkInFiles(media.Location, "cover", "folder"); imageData == nil {
		w := &bytes.Buffer{}
		if err = images.ExtractImageFromMedia(w, media.Location); err == nil {
			imageData = w.Bytes()
		}
	}
	if imageData == nil {
		return nil, nil
	}

	// calculate the sha of the image
	if shaData, err = helpers.SHA(bytes.NewBuffer(imageData)); err != nil {
		logger.Error(err)
		return nil, err
	}
	media.ShaImage = helpers.HashFromSHA(shaData)

	return imageData, nil
}

func FindArtworkInFiles(filename string, matches ...string) ([]byte, error) {
	if len(matches) == 0 {
		return nil, errors.New("missing parameters: matches")
	}
	keys := make(map[string]struct{})
	for _, v := range matches {
		keys[strings.ToLower(v)] = struct{}{}
	}

	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}
	var coverFileName string
	var ext string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		ext = filepath.Ext(f.Name())
		base := filepath.Base(f.Name())
		base = strings.TrimSuffix(base, ext)
		if _, ok := keys[strings.ToLower(base)]; ok {
			coverFileName = f.Name()
			break
		}
	}
	if coverFileName == "" {
		return nil, errors.New("no cover file in directory: " + dir)
	}

	// png and jpg have different paths
	var imageData []byte

	switch strings.ToLower(ext) {
	case ".png", ".jpg":
		if imageData, err = ioutil.ReadFile(filepath.Join(dir, coverFileName)); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("no cover file in directory: " + dir)
	}

	// TODO: allow the images package to deal with jpg as well
	// jpg images need to be converted to png
	switch strings.ToLower(ext) {
	case ".jpg":
		var img image.Image
		if img, err = jpeg.Decode(bytes.NewReader(imageData)); err != nil {
			return nil, err
		}
		w := &bytes.Buffer{}
		if err = png.Encode(w, img); err != nil {
			return nil, err
		}
		return w.Bytes(), nil
	}
	return imageData, nil
}
