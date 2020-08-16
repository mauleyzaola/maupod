package main

import (
	"encoding/json"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/broker"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

// handlerPositionPercentChange may come as JSON format
func (m *MsgHandler) handlerPositionPercentChange(msg *nats.Msg) {
	var input pb.SocketTrackPositionChangeInput
	err := json.Unmarshal(msg.Data, &input)
	if err != nil {
		log.Println(err)
		return
	}

	if err = m.ipc.SeekAbsolute(input.Percent); err != nil {
		log.Println(err)
		return
	}

	// need to trigger another track play event, so UI redraws the spectrum
	if err = broker.PublishBrokerJSON(m.base.NATS(), pb.Message_MESSAGE_SOCKET_PLAY_TRACK, &pb.PlayTrackInput{Media: input.Media}); err != nil {
		log.Println(err)
	}
}
