package pkg

import (
	"log"
	"os"
	"time"

	"github.com/DexterLB/mpvipc"
)

type CommandEnum int

const defaultStartupSecs = 1

const (
	cmdLoadfile CommandEnum = iota
	cmdPause
	cmdSeekOffset
	cmdSeekExact
	cmdVolume
	cmdSpeed
)

func (c *CommandEnum) values() []string {
	return []string{"loadfile", "pause", "seek", "seek", "volume", "speed"}
}

func (c *CommandEnum) String() string {
	return c.values()[*c]
}

type MpvWrapper struct {
	process  *os.Process
	ipc      *mpvipc.Connection
	isPaused bool
}

func NewMPV(filename string) (*MpvWrapper, error) {
	process, err := MPVStart(filename)
	if err != nil {
		return nil, err
	}
	wrapper := &MpvWrapper{
		process:  process,
		ipc:      mpvipc.NewConnection(socketFile),
		isPaused: true,
	}
	time.Sleep(time.Duration(defaultStartupSecs))
	return wrapper, nil
}

func (m *MpvWrapper) Play(filename string) error {
	cmd := cmdLoadfile
	log.Println(cmd.String(), filename)
	_, err := m.ipc.Call(cmd.String(), filename)
	return err
}

func (m *MpvWrapper) PauseToggle() error {
	cmd := cmdPause
	m.isPaused = !m.isPaused
	log.Println(cmd.String(), m.isPaused)
	if err := m.ipc.Set(cmd.String(), m.isPaused); err != nil {
		return err
	}
	return nil
}

func (m *MpvWrapper) IsPaused() bool {
	return m.isPaused
}

func (m *MpvWrapper) Terminate() error {
	_ = m.ipc.Close()
	return m.process.Kill()
}

func (m *MpvWrapper) Seek(secs int) error {
	cmd := cmdSeekOffset
	_, err := m.ipc.Call(cmd.String(), secs, "relative")
	return err
}

func (m *MpvWrapper) SeekExact(secs int) error {
	cmd := cmdSeekExact
	_, err := m.ipc.Call(cmd.String(), secs, "exact")
	return err
}

func (m *MpvWrapper) Volume(v int) error {
	cmd := cmdVolume
	return m.ipc.Set(cmd.String(), v)
}

func (m *MpvWrapper) Speed(v float64) error {
	cmd := cmdSpeed
	return m.ipc.Set(cmd.String(), v)
}
