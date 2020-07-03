package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/pkg/images"

	"github.com/nats-io/nats.go"
)

const thumbnailDir = "thumbnail"

func (m *MsgHandler) handlerArtworkExtract(msg *nats.Msg) {
	var err error
	var input pb.ArtworkExtractInput
	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	//we need to update the database when this function exits, one way or another
	defer func() {
		var payload []byte
		if err != nil {
			return
		}
		if payload, err = helpers.ProtoMarshal(&pb.ArtworkUpdateInput{Media: input.Media}); err != nil {
			log.Println(err)
			return
		}
		if err = m.base.NATS().Publish(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_UPDATE_ARTWORK)), payload); err != nil {
			log.Println(err)
			return
		}
	}()

	var imageData []byte
	var data []byte

	input.Media.LastImageScan = helpers.TimeToTs(helpers.Now())

	// check if artwork already exist for the same album
	var artworkPath = artworkFullPath(m.config, input.Media)
	if _, err = os.Stat(artworkPath); err == nil {
		var file *os.File
		if file, err = os.Open(artworkPath); err != nil {
			log.Println(err)
			return
		}
		if data, err = helpers.SHA(file); err != nil {
			log.Println(err)
			return
		}
		if err = file.Close(); err != nil {
			log.Println(err)
			return
		}
		input.Media.ShaImage = helpers.HashFromSHA(data)
		return
	}

	// check for the image in the same directory of the audio file
	if imageData, err = extractMediaFromDirectory(input.Media.Location); err != nil {
		// check for image in the audio file
		if imageData, err = extractMediaFromAudioFile(input.Media.Location); err != nil {
			return
		}
	}

	// TODO: this is a fucking mess
	// 1. use mediainfo for getting image size
	// 2. use imagemagick for resize and png conversion
	// convert cover.jpg -resize 300x300 cover.png :D

	// no image available, exit
	if imageData == nil {
		return
	}

	// check shape and size are valid
	if err = imageValidSize(data, int(m.config.ArtworkBigSize)); err != nil {
		log.Println(err)
		return
	}

	if data, err = helpers.SHA(bytes.NewBuffer(imageData)); err != nil {
		log.Println(err)
		return
	}
	if err = imageWriteArtwork(imageData, artworkPath, int(m.config.ArtworkBigSize)); err != nil {
		log.Println(err)
		return
	}
	input.Media.ShaImage = helpers.HashFromSHA(data)
}

func artworkFullPath(config *pb.Configuration, media *pb.Media) string {
	var dir = config.ArtworkStore.Location
	var imageLocation = rules.ArtworkFileName(media)
	return filepath.Join(dir, imageLocation)
}

func imageWriteArtwork(data []byte, artworkPath string, imageSize int) error {
	var bigSize = imageSize
	err := images.ImageResize(bytes.NewBuffer(data), artworkPath, bigSize, bigSize)
	if err != nil {
		return err
	}
	return nil
}

func imageValidSize(data []byte, minWidth int) error {
	x, y, err := images.Size(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	// check symmetry
	if x != y {
		return fmt.Errorf("invalid image shape: %dx%d", x, y)
	}

	// check min width
	if x < minWidth {
		return errors.New("image too small")
	}
	return nil
}

// picks the image from the audio file
func extractMediaFromAudioFile(location string) ([]byte, error) {
	w := &bytes.Buffer{}
	if err := images.ExtractImageFromMedia(w, location); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

// extractMediaFromDirectory picks first image from directory
func extractMediaFromDirectory(location string) ([]byte, error) {
	const (
		pngExt = ".png"
		jpgExt = ".jpg"
	)

	dir := filepath.Dir(location)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var validImageFiles = helpers.StringSlice([]string{pngExt, jpgExt})
	var matchedFile os.FileInfo
	var ext string
	for _, v := range files {
		ext = filepath.Ext(v.Name())
		ext = strings.ToLower(ext)
		if !validImageFiles.ContainsAny(ext) {
			continue
		}
		matchedFile = v
		break
	}
	if matchedFile == nil {
		return nil, errors.New("no image file found in dir: " + dir)
	}

	var imageData []byte
	if imageData, err = ioutil.ReadFile(filepath.Join(dir, matchedFile.Name())); err != nil {
		return nil, err
	}

	return imageData, nil
}
