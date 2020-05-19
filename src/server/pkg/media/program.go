package media

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

const (
	mediaInfoDateFormat = "UTC 2006-01-02 15:04:05"
)

// MediaInfoFromFiles returns a MediaInfo slice
// params can be either one file, many files or even many paths, not necessary
// pointing to specific audio files, but directories that contain audio files within
func MediaInfoFromFiles(filename string) (*MediaInfo, error) {
	const mediaInfoProgram = "mediainfo"
	if !programExists(mediaInfoProgram) {
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

	return MediaParser(output.Bytes())
}

func programExists(programName string) bool {
	_, err := exec.LookPath(programName)
	return err == nil
}

func MediaInfoWithSHA(filename string, fn func(info *MediaInfo, id string)) error {
	info, err := MediaInfoFromFiles(filename)
	if err != nil {
		return err
	}
	if fn != nil {
		var data []byte
		var hash []byte
		if data, err = ioutil.ReadFile(filename); err != nil {
			return err
		}
		buffer := bytes.NewBuffer(data)
		if hash, err = helpers.SHA(buffer); err != nil {
			return err
		}
		fn(info, fmt.Sprintf("%x", string(hash)))
	}
	return nil
}
