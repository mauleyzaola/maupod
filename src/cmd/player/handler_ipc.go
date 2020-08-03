package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerIPC(msg *nats.Msg) {
	var err error
	var input pb.IPCInput

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	log.Println("command: ", input.Command)

	var filename string

	// check ipc has been initialized
	if val := input.Media.Location; val != "" {
		filename = paths.FullPath(val)
		if err = m.InitializeIPC(filename); err != nil {
			m.base.Logger().Error(err)
			return
		}
	}
	input.Media.Location = filename

	// check the file exists before emit the event to mpv https://github.com/mauleyzaola/maupod/issues/75
	if _, err = os.Stat(input.Media.Location); err != nil {
		log.Println(err)
		return
	}

	switch input.Command {
	case pb.Message_IPC_PLAY:
		if err = m.ipc.Load(input.Media); err != nil {
			m.base.Logger().Error(err)
			return
		}
		if err = m.ipc.Play(); err != nil {
			m.base.Logger().Error(err)
			return
		}
	case pb.Message_IPC_PAUSE:
		if err = m.ipc.PauseToggle(); err != nil {
			m.base.Logger().Error(err)
			return
		}
	case pb.Message_IPC_LOAD:
		if err = m.ipc.Load(input.Media); err != nil {
			m.base.Logger().Error(err)
			return
		}
	case pb.Message_IPC_SKIP:
		if input.Media == nil {
			return
		}
		m.ipc.Skip()
	case pb.Message_IPC_VOLUME:
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
