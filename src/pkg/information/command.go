package information

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
)

const MediaInfoProgram = "mediainfo"

func MediaInfoFromFile(filename string) (*bytes.Buffer, error) {
	if !helpers.ProgramExists(MediaInfoProgram) {
		return nil, fmt.Errorf("could not find program: %s in path", MediaInfoProgram)
	}
	var p = []string{
		"-f",
	}
	p = append(p, filename)
	cmd := exec.Command(MediaInfoProgram, p...)
	output := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = errOutput
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%s %s : %v", output.String(), errOutput.String(), err)
	}
	return output, nil
}
