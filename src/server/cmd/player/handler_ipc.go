package main

import (
	"fmt"
	"os"
	"strings"

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

	var filename string

	// check ipc has been initialized
	if val := input.Media.Location; val != "" {
		filename = convertToLocalPath(val)
		if err = m.InitializeIPC(filename); err != nil {
			m.base.Logger().Error(err)
			return
		}
	}

	switch input.Command {
	case pb.IPCCommand_IPC_PLAY:
		if err = m.ipc.Play(); err != nil {
			output.Error = err.Error()
			output.Ok = false
			return
		}
	case pb.IPCCommand_IPC_PAUSE:
		if err = m.ipc.Pause(); err != nil {
			output.Error = err.Error()
			output.Ok = false
			return
		}
	case pb.IPCCommand_IPC_LOAD:
		if err = m.ipc.Load(filename); err != nil {
			output.Error = err.Error()
			output.Ok = false
			return
		}
	default:
		output.Error = fmt.Sprintf("unsupported command: %v", input.Command)
		output.Ok = false
		return
	}

	output.Ok = true

	return
}

// TODO: fix this mess with the volume name vs local path
func convertToLocalPath(filename string) string {
	val := os.Getenv("MEDIA_STORE")
	if val == "" {
		return filename
	}
	return strings.Replace(filename, "/music-store", val, -1)
}
