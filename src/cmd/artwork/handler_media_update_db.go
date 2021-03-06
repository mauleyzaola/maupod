package main

import (
	"log"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/taggers"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerTagUpdate(msg *nats.Msg) {
	var err error
	var input protos.MediaInfoInput

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	tagger, err := taggers.TaggerFactory(input.Media.Location)
	if err != nil {
		log.Println(err)
		return
	}

	if err = tagger.Tag(input.Media, input.Media.Location); err != nil {
		log.Println(err)
		return
	}
	return
}
