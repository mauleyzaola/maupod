package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerIPC(msg *nats.Msg) {
	var err error
	var input pb.IPCInput

	if err = json.Unmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		return
	}

	log.Println("command: ", input.Command)

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
	input.Media.Location = filename

	switch input.Command {
	case pb.IPCCommand_IPC_PLAY:
		if err = m.ipc.Load(input.Media); err != nil {
			m.base.Logger().Error(err)
			return
		}
		if err = m.ipc.Play(); err != nil {
			m.base.Logger().Error(err)
			return
		}
	case pb.IPCCommand_IPC_PAUSE:
		if err = m.ipc.PauseToggle(); err != nil {
			m.base.Logger().Error(err)
			return
		}
	case pb.IPCCommand_IPC_LOAD:
		if err = m.ipc.Load(input.Media); err != nil {
			m.base.Logger().Error(err)
			return
		}
	case pb.IPCCommand_IPC_VOLUME:
		val, err := strconv.ParseInt(input.Value, 10, 64)
		if err != nil {
			m.base.Logger().Error(err)
			return
		}
		if err = m.ipc.Volume(int(val)); err != nil {
			m.base.Logger().Error(err)
			return
		}
	default:
		err = fmt.Errorf("unsupported command: %v", input.Command)
		m.base.Logger().Error(err)
		return
	}

	return
}

// TODO: fix this mess with the volume name vs local path, maybe we should store file paths relative to the MEDIA_STORE value
func convertToLocalPath(filename string) string {
	val := os.Getenv("MEDIA_STORE")
	if val == "" {
		return filename
	}
	return strings.Replace(filename, "/music-store", val, -1)
}
