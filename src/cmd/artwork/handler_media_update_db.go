package main

import (
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/taggers"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerTagUpdate(msg *nats.Msg) {
	var err error
	var input pb.MediaInfoInput

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		return
	}

	tagger, err := taggers.TaggerFactory(input.Media.Location)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}

	if err = tagger.Tag(input.Media, input.Media.Location); err != nil {
		m.base.Logger().Error(err)
		return
	}
	return
}
