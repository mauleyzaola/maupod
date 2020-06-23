package pkg

import (
	"errors"
	"log"
	"os"

	"github.com/DexterLB/mpvipc"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
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
	lastMedia  *pb.Media
	control    *PlayerControl
}

func NewIPC(processor MPVProcessor, control *PlayerControl) (*IPC, error) {
	if processor == nil {
		return nil, errors.New("missing parameter: processor")
	}

	connection := mpvipc.NewConnection(processor.SocketFileName())

	ipc := &IPC{
		connection: connection,
		isPaused:   true,
		processor:  processor,
		control:    control,
	}
	// configure which events will be listening to mpv actions
	ipc.listeners = map[pb.Message]*EventListener{
		//pb.Message_MESSAGE_MPV_FILENAME: {
		//	eventName: "filename",
		//	trigger:   ipc.triggerFilename,
		//},
		//pb.Message_MESSAGE_MPV_STREAM_POS: {
		//	eventName: "stream-pos",
		//	trigger:   ipc.triggerStreamPos,
		//},
		//pb.Message_MESSAGE_MPV_STREAM_END: {
		//	eventName: "stream-end",
		//	trigger:   ipc.triggerStreamEnd,
		//},
		pb.Message_MESSAGE_MPV_PERCENT_POS: {
			eventName: "percent-pos",
			trigger:   ipc.triggerPercentPos,
		},
		pb.Message_MESSAGE_MPV_TIME_POS: {
			eventName: "time-pos",
			trigger:   ipc.triggerTimePos,
		},
		pb.Message_MESSAGE_MPV_TIME_REMAINING: {
			eventName: "time-remaining",
			trigger:   ipc.triggerTimeRemaining,
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

func (m *IPC) Load(media *pb.Media) error {
	var filename = media.Location
	err := m.restart(filename)
	if err != nil {
		return err
	}
	cmd := cmdLoadfile
	if _, err = os.Stat(filename); err != nil {
		return err
	}
	m.lastMedia = media
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
	_, err := m.connection.Call(cmd.String(), secs, "relative")
	return err
}

func (m *IPC) SeekExact(secs int) error {
	cmd := cmdSeekExact
	_, err := m.connection.Call(cmd.String(), secs, "exact")
	return err
}

func (m *IPC) Volume(v int) error {
	cmd := cmdVolume
	return m.connection.Set(cmd.String(), v)
}

func (m *IPC) Speed(v float64) error {
	cmd := cmdSpeed
	return m.connection.Set(cmd.String(), v)
}

func (m *IPC) Play() error {
	if err := m.pause(false); err != nil {
		return err
	}
	m.control.OnSongStarted(m.lastMedia)
	return nil
}
