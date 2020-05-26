package media

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

// RunMediaInfo returns a MediaInfo slice
// params can be either one file, many files or even many paths, not necessary
// pointing to specific audio files, but directories that contain audio files within
func RunMediaInfo(filename string) (*MediaInfo, error) {
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
		return nil, err
	}

	return MediaParser(output)
}

// TODO: implement a unique field parser to discover more fields in the future from mediainfo output
//func MediaInfoFieldFinder(filename string) ([]string, error) {
//	return nil, errors.New("not implemented")
//}
