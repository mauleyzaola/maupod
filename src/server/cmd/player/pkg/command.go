package pkg

import (
	"errors"
	"fmt"
	"log"
	"os/exec"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

const mpvProgram = "mpv"

// MPVCommand forks a process passing the default parameters and plays a track
func MPVCommand(container MPVSocketFileContainer, trackPath string) (*exec.Cmd, error) {
	if container == nil {
		return nil, errors.New("missing parameter: container")
	}
	if !helpers.ProgramExists(mpvProgram) {
		return nil, errors.New("could not find processor executable on system")
	}
	var pars = []string{"--no-video", fmt.Sprintf("--input-ipc-server=%s", container.SocketFileName()), "--quiet", "--pause", trackPath}
	cmd := exec.Command(mpvProgram, pars...)
	log.Println("created  processor command with params: ", pars)
	return cmd, nil
}
