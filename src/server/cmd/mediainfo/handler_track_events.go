package main

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerTrackPlayCountIncrease(msg *nats.Msg) {
	var input pb.TrackPlayedInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	m.base.Logger().Info("handlerTrackPlayCountIncrease")
	m.base.Logger().Info(input.String())
}

func (m *MsgHandler) handlerTrackSkipped(msg *nats.Msg) {
	var input pb.TrackSkippedInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	m.base.Logger().Info("handlerTrackSkipped" + input.String())
}
