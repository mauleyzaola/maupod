package main

import (
	"encoding/json"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
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
	duration, err := rules.MediaPercentToSeconds(input.Media, input.Percent)
	if err != nil {
		log.Println(err)
		return
	}
	if err = m.ipc.SeekExact(int(duration.Seconds())); err != nil {
		log.Println(err)
		return
	}
}
