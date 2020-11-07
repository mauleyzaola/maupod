package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerMediaSpectrumGenerate(msg *nats.Msg) {
	var input protos.SpectrumGenerateInput
	var output protos.SpectrumGenerateOutput
	var err error

	defer func() {
		if msg.Reply == "" {
			return
		}
		data, err := helpers.ProtoMarshal(&output)
		if err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
			return
		}
	}()

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		output.Error = err.Error()
		return
	}
	w := &bytes.Buffer{}
	if err = generateSpectrum(w, paths.MediaFullPathAudioFile(input.Media.Location), int(input.Width), int(input.Height)); err != nil {
		output.Error = err.Error()
		return
	}
	output.Media = input.Media
	output.Data = w.Bytes()
	return
}

// TODO: allow to set the color
// TODO: allow to set the width and height
func generateSpectrum(w io.Writer, filename string, width, heigth int) error {
	const ffmpegProgram = "ffmpeg"
	if !helpers.ProgramExists(ffmpegProgram) {
		return fmt.Errorf("could not find program: %s in path", ffmpegProgram)
	}
	destination := filepath.Join(os.TempDir(), helpers.NewUUID()) + ".png"

	var p = []string{
		"-i",
		filename,
		"-lavfi",
		//"showwavespic=split_channels=0:s=1920x108:colors=48c3e8",
		fmt.Sprintf("showwavespic=split_channels=0:s=%dx%d:colors=48c3e8", width, heigth),
		destination,
	}
	cmd := exec.Command(ffmpegProgram, p...)
	output := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = errOutput
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s %s : %v", output.String(), errOutput.String(), err)
	}

	file, err := os.Open(destination)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
			return
		}
		_ = os.Remove(destination)
	}()
	if _, err = io.Copy(w, file); err != nil {
		return err
	}
	return nil
}
