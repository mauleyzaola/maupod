package main

import (
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

	// check ipc has been initialized always
	if val := input.Media.Location; val != "" {
		if err = m.InitializeIPC(val); err != nil {
			m.base.Logger().Error(err)
			return
		}
	}
	m.base.Logger().Info("received ipc message: " + input.String())

	return
}
