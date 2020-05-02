package media

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

// Mediainfo returns a MediaInfo slice
// params can be either one file, many files or even many paths, not necessary
// pointing to specific audio files, but directories that contain audio files within
func Mediainfo(params ...string) ([]MediaInfo, error) {
	const mediaInfoProgram = "mediainfo"
	if !programExists(mediaInfoProgram) {
		return nil, fmt.Errorf("could not find program: %s in path", mediaInfoProgram)
	}
	var p = []string{
		"--Output=JSON",
	}
	p = append(p, params...)
	cmd := exec.Command(mediaInfoProgram, p...)
	output := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = errOutput
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// mediainfo when using JSON format, can return one object or an array, need to check both options
	var payload MediaInfo
	var payloads []MediaInfo
	if err = json.Unmarshal(output.Bytes(), &payload); err == nil {
		return []MediaInfo{payload}, nil
	}
	if err = json.Unmarshal(output.Bytes(), &payloads); err != nil {
		return nil, err
	}

	return payloads, nil
}

func programExists(programName string) bool {
	_, err := exec.LookPath(programName)
	return err == nil
}
