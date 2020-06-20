package pkg

import (
	"errors"
	"os"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

const mpvProgram = "mpv"

// MPVStart forks a process passing the default parameters and plays a track
func MPVStart(trackPath string) (*os.Process, error) {
	if !helpers.ProgramExists(mpvProgram) {
		return nil, errors.New("could not find mpv executable on system")
	}

	// with these parameters mpv will need to issue a pause set to false to start playing
	// giving the freedom of choice to load another file :D
	// conn.Set("pause", false)
	var pars = []string{"--no-video", "--input-unix-socket=/tmp/mpv_socket", "--quiet", "--pause", trackPath}

	cmd := exec.Command(mpvProgram, pars...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd.Process, nil
}
