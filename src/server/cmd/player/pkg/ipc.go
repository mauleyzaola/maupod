package pkg

import (
	"errors"
	"log"
	"os"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"

	"github.com/DexterLB/mpvipc"
)

type DispatcherFunc func(v interface{})

type EventListener struct {
	eventName string
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
	connection *mpvipc.Connection
	isPaused   bool
	processor  MPVProcessor
	listeners  map[pb.Message]*EventListener
}

func NewIPC(processor MPVProcessor) (*IPC, error) {
	if processor == nil {
		return nil, errors.New("missing parameter: processor")
	}

	connection := mpvipc.NewConnection(processor.SocketFileName())

	ipc := &IPC{
		connection: connection,
		isPaused:   true,
		processor:  processor,
		// configure which events will be listening to mpv actions
		listeners: map[pb.Message]*EventListener{
			pb.Message_MESSAGE_MPV_FILENAME: {
				eventName: "filename",
				trigger:   triggerFilename,
			},
			pb.Message_MESSAGE_MPV_STREAM_POS: {
				eventName: "stream-pos",
				trigger:   triggerStreamPos,
			},
			pb.Message_MESSAGE_MPV_STREAM_END: {
				eventName: "stream-end",
				trigger:   triggerStreamEnd,
			},
			pb.Message_MESSAGE_MPV_PERCENT_POS: {
				eventName: "percent-pos",
				trigger:   triggerPercentPos,
			},
			pb.Message_MESSAGE_MPV_TIME_POS: {
				eventName: "time-pos",
				trigger:   triggerTimePos,
			},
			pb.Message_MESSAGE_MPV_TIME_REMAINING: {
				eventName: "time-remaining",
				trigger:   triggerTimeRemaining,
			},
		},
	}

	return ipc, nil
}

func (m *IPC) configureListeners() error {
	const observeCommand = "observe_property"
	if len(m.listeners) == 0 {
		return nil
	}

	// properties for monitoring what is exactly doing mpv at runtime
	events, stopListening := m.connection.NewEventListener()
	go func() {
		log.Println("connection will close now")
		m.connection.WaitUntilClosed()
		stopListening <- struct{}{}
	}()

	var err error
	for k, v := range m.listeners {
		log.Printf("registering listener: %v %v\n", k, v.eventName)
		if _, err = m.connection.Call(observeCommand, k, v.eventName); err != nil {
			return err
		}
	}

	go func() {
		for v := range events {
			dispatcher, ok := m.listeners[pb.Message(v.ID)]
			if !ok {
				continue
			}
			dispatcher.trigger(v.Data)
		}
	}()
	return nil
}

// restart checks if both mpv and connection are open and running
// if they are not, it will start/open them
func (m *IPC) restart(filename string) error {
	// TODO: improve this workflow, for now we can start and stop mpv process at will
	var err error
	if !m.processor.IsRunning() {
		if err = m.processor.Start(filename); err != nil {
			return err
		}
	}
	if m.connection.IsClosed() {
		if err = m.connection.Open(); err != nil {
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
	_, err = m.connection.Call(cmd.String(), filename)
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
	if err := m.connection.Set(cmd.String(), v); err != nil {
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
		_ = m.connection.Close()
	}()
	if m.processor != nil {
		return m.processor.Close()
	}
	return nil
}

func (m *IPC) Seek(secs int) error {
	cmd := cmdSeekOffset
	log.Println(cmd.String(), secs)
	_, err := m.connection.Call(cmd.String(), secs, "relative")
	return err
}

func (m *IPC) SeekExact(secs int) error {
	cmd := cmdSeekExact
	log.Println(cmd.String(), "exact", secs)
	_, err := m.connection.Call(cmd.String(), secs, "exact")
	return err
}

func (m *IPC) Volume(v int) error {
	cmd := cmdVolume
	log.Println(cmd.String(), v)
	return m.connection.Set(cmd.String(), v)
}

func (m *IPC) Speed(v float64) error {
	cmd := cmdSpeed
	return m.connection.Set(cmd.String(), v)
}

func (m *IPC) Play() error {
	cmd := cmdPlay
	log.Println(cmd.String())
	if err := m.pause(false); err != nil {
		return err
	}
	return nil
}
