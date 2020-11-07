package main

import (
	"log"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerSHAScan(msg *nats.Msg) {
	var err error
	var input protos.SHAScanInput

	defer func() {
		// TODO: handle err != nil
	}()

	// read SHA from media
	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}
	var location = paths.LocationPath(input.Media.Location)
	var fullPath = paths.MediaFullPathAudioFile(location)
	sha, err := helpers.SHAFromFile(fullPath)
	if err != nil {
		log.Println(err)
		return
	}
	// if sha different than provided SHA, send an update SHA message
	if err = broker.PublishMediaSHAUpdate(m.base.NATS(), &protos.MediaUpdateSHAInput{
		Media:  input.Media,
		OldSHA: input.Media.Sha,
		NewSHA: sha,
	}); err != nil {
		log.Println(err)
		return
	}
}
