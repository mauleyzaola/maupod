package main

import (
	"encoding/json"
	"github.com/mauleyzaola/maupod/src/protos"
	"log"

	"github.com/nats-io/nats.go"
)

// handlerPositionPercentChange may come as JSON format
func (m *MsgHandler) handlerPositionPercentChange(msg *nats.Msg) {
	var input protos.SocketTrackPositionChangeInput
	err := json.Unmarshal(msg.Data, &input)
	if err != nil {
		log.Println(err)
		return
	}

	if err = m.ipc.SeekAbsolute(input.Percent); err != nil {
		log.Println(err)
		return
	}
}
