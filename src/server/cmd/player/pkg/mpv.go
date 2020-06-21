package pkg

import (
	"log"
	"os"
)

type MPVProcessCloser interface {
	Close() error
}

type MPVProcess struct {
	process *os.Process
}

func NewMPVProcess(filename string) (MPVProcessCloser, error) {
	cmd, err := MPVCommand(filename)
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	log.Println("pid: ", cmd.Process.Pid)

	// give mpvProcessCloser some time to start up
	sleep(defaultStartupSecs)

	return &MPVProcess{
		process: cmd.Process,
	}, nil
}

func (mpv *MPVProcess) Close() error {
	return mpv.process.Kill()
}
