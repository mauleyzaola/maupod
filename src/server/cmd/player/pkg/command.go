package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

const mpvProgram = "mpv"

var socketFile = filepath.Join(os.TempDir(), "mpv_socket")

// MPVCommand forks a process passing the default parameters and plays a track
func MPVCommand(trackPath string) (*exec.Cmd, error) {
	if !helpers.ProgramExists(mpvProgram) {
		return nil, errors.New("could not find mpvProcessCloser executable on system")
	}
	var pars = []string{"--no-video", fmt.Sprintf("--input-ipc-server=%s", socketFile), "--quiet", "--pause", trackPath}
	cmd := exec.Command(mpvProgram, pars...)
	log.Println("created  mpvProcessCloser command with params: ", pars)
	return cmd, nil
}
