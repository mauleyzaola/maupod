package pkg

import (
	"log"

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
	cmdPlay
)

func (c *CommandEnum) values() []string {
	return []string{"loadfile", "pause", "seek", "seek", "volume", "speed", "play"}
}

func (c *CommandEnum) String() string {
	return c.values()[*c]
}

type IPC struct {
	ipc       *mpvipc.Connection
	isPaused  bool
	processor MPVProcessor
}

func NewIPC(processor MPVProcessor) (*IPC, error) {
	wrapper := &IPC{
		isPaused:  true,
		processor: processor,
	}

	wrapper.ipc = mpvipc.NewConnection(processor.SocketFileName())
	if err := wrapper.ipc.Open(); err != nil {
		return nil, err
	}
	log.Println("opened successfully ipc wrapper for processor socket: ", processor.SocketFileName())
	return wrapper, nil
}

func (m *IPC) Load(filename string) error {
	cmd := cmdLoadfile
	log.Println(cmd.String(), filename)
	_, err := m.ipc.Call(cmd.String(), filename)
	return err
}

func (m *IPC) PauseToggle() error {
	m.isPaused = !m.isPaused
	if err := m.pause(m.isPaused); err != nil {
		return err
	}
	return nil
}

func (m *IPC) pause(v bool) error {
	cmd := cmdPause
	log.Println(cmd.String(), v)
	if err := m.ipc.Set(cmd.String(), v); err != nil {
		return err
	}
	m.isPaused = v
	return nil
}

func (m *IPC) Pause() error {
	return m.pause(true)
}

func (m *IPC) IsPaused() bool {
	return m.isPaused
}

func (m *IPC) Terminate() error {
	log.Println("received termination signal")
	_ = m.Pause()
	sleep(defaultStartupSecs)
	defer func() {
		_ = m.ipc.Close()
	}()
	if m.processor != nil {
		return m.processor.Close()
	}
	return nil
}

func (m *IPC) Seek(secs int) error {
	cmd := cmdSeekOffset
	_, err := m.ipc.Call(cmd.String(), secs, "relative")
	return err
}

func (m *IPC) SeekExact(secs int) error {
	cmd := cmdSeekExact
	_, err := m.ipc.Call(cmd.String(), secs, "exact")
	return err
}

func (m *IPC) Volume(v int) error {
	cmd := cmdVolume
	return m.ipc.Set(cmd.String(), v)
}

func (m *IPC) Speed(v float64) error {
	cmd := cmdSpeed
	return m.ipc.Set(cmd.String(), v)
}

func (m *IPC) Play() error {
	cmd := cmdPlay
	log.Println(cmd.String())
	if err := m.pause(false); err != nil {
		return err
	}
	return nil
}
