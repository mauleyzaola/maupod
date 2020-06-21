package main

import (
	"encoding/json"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerIPC(msg *nats.Msg) {

	var err error
	var input pb.IPCInput

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		return
	}
	data, err := json.MarshalIndent(&input, "", "  ")
	if err != nil {
		m.base.Logger().Error(err)
		return
	}
	m.base.Logger().Info("received ipc message: ")
	m.base.Logger().Info(string(data))

	output := &pb.IPCOutput{
		Ok:    false,
		Error: "",
	}

	defer func() {
		var localErr error
		var data []byte

		if data, localErr = helpers.ProtoMarshal(output); localErr != nil {
			m.base.Logger().Error(localErr)
			return
		}
		if localErr = msg.Respond(data); localErr != nil {
			m.base.Logger().Error(localErr)
			return
		}
	}()

	// TODO: validate media is ok

	// check ipc has been initialized always
	if val := input.Media.Location; val != "" {
		if err = m.InitializeIPC(val); err != nil {
			m.base.Logger().Error(err)
			return
		}
	}

	output.Ok = true

	return
}
