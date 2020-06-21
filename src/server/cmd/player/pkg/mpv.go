package pkg

import (
	"log"
	"os"
	"path/filepath"
)

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

func NewMPVProcess(filename string) (MPVProcessor, error) {
	mpv := &MPVProcess{
		socketFile: filepath.Join(os.TempDir(), "mpv_socket"),
	}
	cmd, err := MPVCommand(mpv, filename)
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	log.Println("pid: ", cmd.Process.Pid)
	mpv.process = cmd.Process

	// give processor some time to start up
	sleep(defaultStartupSecs)

	return mpv, nil
}

func (mpv *MPVProcess) Close() error {
	return mpv.process.Kill()
}

func (mpv *MPVProcess) SocketFileName() string {
	return mpv.socketFile
}
