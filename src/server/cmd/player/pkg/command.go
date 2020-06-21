package pkg

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

const mpvProgram = "mpv"

var socketFile = filepath.Join(os.TempDir(), "mpv_socket")

// MPVStart forks a process passing the default parameters and plays a track
func MPVStart(trackPath string) (*os.Process, error) {
	if !helpers.ProgramExists(mpvProgram) {
		return nil, errors.New("could not find mpv executable on system")
	}

	// with these parameters mpv will need to issue a pause set to false to start playing
	// giving the freedom of choice to load another file :D
	// conn.Set("pause", false)
	var pars = []string{"--no-video", fmt.Sprintf("--input-unix-socket=%s", socketFile), "--quiet", "--pause", trackPath}

	cmd := exec.Command(mpvProgram, pars...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd.Process, nil
}
