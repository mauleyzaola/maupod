package pkg

import (
	"errors"
	"log"
	"os"

	"github.com/DexterLB/mpvipc"
)

type DispatcherFunc func(v interface{})

type EventListener struct {
	id        uint
	eventName string
	args      string
	trigger   DispatcherFunc
}

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
	listeners map[uint]*EventListener
}

func NewIPC(processor MPVProcessor) (*IPC, error) {
	if processor == nil {
		return nil, errors.New("missing parameter: processor")
	}

	ipc := mpvipc.NewConnection(processor.SocketFileName())

	ip := &IPC{
		ipc:       ipc,
		isPaused:  true,
		processor: processor,
		listeners: map[uint]*EventListener{}, // TODO: configure which events will be listening to
	}

	return ip, nil
}

func (m *IPC) configureListeners() error {
	if len(m.listeners) == 0 {
		return nil
	}

	// properties for monitoring what is exactly doing mpv at runtime
	events, stopListening := m.ipc.NewEventListener()
	go func() {
		log.Println("ipc will close now")
		m.ipc.WaitUntilClosed()
		stopListening <- struct{}{}
	}()

	var err error
	for k, v := range m.listeners {
		if _, err = m.ipc.Call(v.eventName, k, v.args); err != nil {
			return err
		}
	}

	go func() {
		for v := range events {
			dispatcher, ok := m.listeners[v.ID]
			if !ok {
				continue
			}
			dispatcher.trigger(v.Data)
		}
	}()
	return nil
}

// restart checks if both mpv and ipc are open and running
// if they are not, it will start/open them
func (m *IPC) restart(filename string) error {
	// TODO: improve this workflow, for now we can start and stop mpv process at will
	var err error
	if !m.processor.IsRunning() {
		if err = m.processor.Start(filename); err != nil {
			return err
		}
	}
	if m.ipc.IsClosed() {
		if err = m.ipc.Open(); err != nil {
			return err
		}
		if err = m.configureListeners(); err != nil {
			return err
		}
	}
	return nil
}

func (m *IPC) Load(filename string) error {
	err := m.restart(filename)
	if err != nil {
		return err
	}
	cmd := cmdLoadfile
	log.Println(cmd.String(), filename)
	if _, err = os.Stat(filename); err != nil {
		return err
	}
	_, err = m.ipc.Call(cmd.String(), filename)
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
	log.Println(cmd.String(), secs)
	_, err := m.ipc.Call(cmd.String(), secs, "relative")
	return err
}

func (m *IPC) SeekExact(secs int) error {
	cmd := cmdSeekExact
	log.Println(cmd.String(), "exact", secs)
	_, err := m.ipc.Call(cmd.String(), secs, "exact")
	return err
}

func (m *IPC) Volume(v int) error {
	cmd := cmdVolume
	log.Println(cmd.String(), v)
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
