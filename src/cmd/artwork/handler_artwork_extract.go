package main

import (
	"log"

	"github.com/mauleyzaola/maupod/src/cmd/artwork/pkg/artworks"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

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
	// try to scan from directory first
	if err = artworks.ExtractFromCoverFile(m.base.NATS(), m.config, media); err != nil {
		log.Println(err)
		// if not possible, try to scan from audio files
		if err = artworks.ExtractWithinAudioFile(m.base.NATS(), m.config, media); err != nil {
			log.Println(err)
			return
		}
	}
	return
}
