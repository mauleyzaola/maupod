package information

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
)

func MediaInfoFromFile(filename string) (*bytes.Buffer, error) {
	const mediaInfoProgram = "mediainfo"
	if !helpers.ProgramExists(mediaInfoProgram) {
		return nil, fmt.Errorf("could not find program: %s in path", mediaInfoProgram)
	}
	var p = []string{
		"-f",
	}
	p = append(p, filename)
	cmd := exec.Command(mediaInfoProgram, p...)
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

// TODO: implement a unique field parser to discover more fields in the future from mediainfo output
//func MediaInfoFieldFinder(filename string) ([]string, error) {
//	return nil, errors.New("not implemented")
//}
