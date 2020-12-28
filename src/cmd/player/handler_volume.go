package main

import (
	"bytes"
	fmt "fmt"
	"log"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerVolumeChange(msg *nats.Msg) {
	input := &protos.VolumeChangeInput{}
	output := &protos.VolumeChangeOutput{}
	defer func() {
		if msg.Reply == "" {
			return
		}
		data, err := helpers.ProtoMarshal(output)
		if err != nil {
			log.Println(err)
			return
		}
		if err := msg.Respond(data); err != nil {
			log.Println(err)
		}
	}()
	if err := helpers.ProtoUnmarshalJSON(msg.Data, input); err != nil {
		log.Println(err)
		log.Println("data: ", string(msg.Data))
		output.Error = err.Error()
		output.Ok = false
		return
	}
	const program = "amixer"
	if !helpers.ProgramExists(program) {
		output.Error = "cannot find amixer binary in path"
		output.Ok = false
		return
	}

	offset := input.Offset
	if offset < 0 || offset > 100 {
		output.Error = "value out of range"
		output.Ok = false
	}

	var value string
	if offset < 0 {
		value = fmt.Sprintf("%d%%-", -offset)
	} else {
		value = fmt.Sprintf("%d%%+", offset)
	}

	var p = []string{
		"-q", "-D", "pulse", "sset", "Master",
		value,
	}
	cmd := exec.Command(program, p...)
	out := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = out
	cmd.Stderr = errOutput
	err := cmd.Run()
	if err != nil {
		output.Error = fmt.Sprintf("%s %s : %v", output.String(), errOutput.String(), err)
		output.Ok = false
		log.Println(output.Error)
		return
	}
	output.Ok = true
	output.Error = ""
}
