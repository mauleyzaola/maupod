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

// mpvCommand forks a process passing the default parameters and plays a track
func mpvCommand(container MPVSocketFileContainer, trackPath string) (*exec.Cmd, error) {
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

type MPVProcessor interface {
	MPVCloser
	MPVSocketFileContainer
}

type MPVCloser interface {
	Close() error
}

type MPVSocketFileContainer interface {
	SocketFileName() string
}

type MPVProcess struct {
	process    *os.Process
	socketFile string
}

func NewMpvProcessor(filename string) (MPVProcessor, error) {
	mpv := &MPVProcess{
		socketFile: filepath.Join(os.TempDir(), "mpv_socket"),
	}
	cmd, err := mpvCommand(mpv, filename)
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	mpv.process = cmd.Process
	log.Println("pid: ", mpv.process.Pid)

	// give processor some time to start up
	sleep(defaultStartupSecs)
	log.Println("mpv started")

	return mpv, nil
}

func (mpv *MPVProcess) Close() error {
	return mpv.process.Kill()
}

func (mpv *MPVProcess) SocketFileName() string {
	return mpv.socketFile
}
