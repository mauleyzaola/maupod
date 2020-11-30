package main

import (
	"encoding/json"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/nats-io/nats.go"
)

// handlerPositionPercentChange may come as JSON format
func (m *MsgHandler) handlerPositionPercentChange(msg *nats.Msg) {
	// TODO: fix this hack
	// msg.Data can come as JSON from node or protobuf from handler
	var input protos.SocketTrackPositionChangeInput
	err := helpers.ProtoUnmarshalJSON(msg.Data, &input)
	if err != nil {
		if err = json.Unmarshal(msg.Data, &input); err != nil {
			log.Println(err)
			return
		}
	}

	if err = m.ipc.SeekAbsolute(input.Percent); err != nil {
		log.Println(err)
		return
	}
}
