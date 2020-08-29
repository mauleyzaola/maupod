package main

import (
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/rules"

	"github.com/mauleyzaola/maupod/src/cmd/artwork/pkg/artworks"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

type extractFn func(conn *nats.Conn, configuration *pb.Configuration, media *pb.Media) error

// handlerArtworkExtract this will only look for image files in the same directory of the audio files
// no scanning of audio files content should be done
func (m *MsgHandler) handlerArtworkExtract(msg *nats.Msg) {
	var err error
	var input pb.ArtworkExtractInput
	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}
	var media = input.Media
	var extractFns = []extractFn{
		artworks.ExtractFromCoverFile,
		artworks.ExtractWithinAudioFile,
	}

	// try to scan using any method
	var ok bool
	for _, fn := range extractFns {
		if err = fn(m.base.NATS(), m.config, media); err != nil {
			log.Println(err)
		} else {
			ok = true
			break
		}
	}
	lastImageScanDate := helpers.TimeToTs(helpers.Now())
	media.LastImageScan = lastImageScanDate
	if ok {
		media.ImageLocation = rules.ArtworkFileName(media)
	} else {
		media.ImageLocation = ""
	}
	if err = artworks.PublishSaveArtworkTrack(m.base.NATS(), media); err != nil {
		log.Println(err)
		return
	}
	if ok {
		log.Println("[INFO] successfully scanned artwork and updated db")
	}

	return
}
