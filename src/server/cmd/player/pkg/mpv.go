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
	var pars = []string{
		"--no-video",
		fmt.Sprintf("--input-ipc-server=%s", container.SocketFileName()),
		"--quiet",
		"--pause",
		"--keep-open=yes",
		trackPath,
	}
	cmd := exec.Command(mpvProgram, pars...)
	log.Println("created  processor command with params: ", pars)
	return cmd, nil
}

type MPVProcessor interface {
	MPVCloser
	MPVRunner
	MPVSocketFileContainer
	MPVStarter
}

type MPVRunner interface {
	IsRunning() bool
}

type MPVStarter interface {
	Start(filename string) error
}

type MPVCloser interface {
	Close() error
}

type MPVSocketFileContainer interface {
	SocketFileName() string
}

type MPVProcess struct {
	process   *os.Process
	isRunning bool
}

func (mpv *MPVProcess) socketfile() string {
	return filepath.Join(os.TempDir(), "mpv_socket")
}

func NewMpvProcessor() (MPVProcessor, error) {
	mpv := &MPVProcess{
		isRunning: false,
	}
	return mpv, nil
}

func (mpv *MPVProcess) Close() error {
	if mpv.process == nil {
		return nil
	}
	return mpv.process.Kill()
}

func (mpv *MPVProcess) SocketFileName() string {
	return mpv.socketfile()
}

func (mpv *MPVProcess) Start(filename string) error {
	if mpv.isRunning {
		if mpv.process != nil {
			_ = mpv.process.Kill()
		}
		mpv.isRunning = false
	}
	cmd, err := mpvCommand(mpv, filename)
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	defer func() {
		mpv.isRunning = true
	}()
	mpv.process = cmd.Process
	log.Println("pid: ", mpv.process.Pid)

	// give processor some time to start up
	sleep(defaultStartupSecs)
	log.Println("mpv started")
	return nil
}

func (mpv *MPVProcess) IsRunning() bool {
	return mpv.isRunning
}
