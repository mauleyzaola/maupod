package images

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
)

func ExtractImageFromMedia(source, target string) error {
	const ffmpeg = "ffmpeg"
	if !helpers.ProgramExists(ffmpeg) {
		return fmt.Errorf("could not find program: %s in path", ffmpeg)
	}
	var p = []string{
		"-i",
		source,
		target,
		"-y",
	}
	cmd := exec.Command(ffmpeg, p...)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(data) + " : " + err.Error())
	}

	return nil
}
