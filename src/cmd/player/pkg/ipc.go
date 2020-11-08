package pkg

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/DexterLB/mpvipc"
)

type DispatcherFunc func(v interface{})

type EventListener struct {
	eventName string
	trigger   DispatcherFunc
}

type CommandEnum int

type PlayedStateFunc func(media *protos.Media, percent float64)

const defaultStartupSecs = 1

const (
	cmdLoadfile CommandEnum = iota
	cmdPause
	cmdSeekOffset
	cmdSeekAbsolute
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
	listeners  map[protos.Message]*EventListener
	lastMedia  *protos.Media
	control    *PlayerControl

	// workaround to avoid a second play event
	lastStartTrackEvent time.Time

	// keep state of the last media played on the last position reported by the IPC
	playedStateFn PlayedStateFunc
}

func NewIPC(processor MPVProcessor, control *PlayerControl, playedStateFn PlayedStateFunc) (*IPC, error) {
	if processor == nil {
		return nil, errors.New("missing parameter: processor")
	}

	connection := mpvipc.NewConnection(processor.SocketFileName())

	ipc := &IPC{
		connection:    connection,
		isPaused:      true,
		processor:     processor,
		control:       control,
		playedStateFn: playedStateFn,
	}
	// configure which events will be listening to mpv actions
	ipc.listeners = map[protos.Message]*EventListener{
		protos.Message_MESSAGE_MPV_PERCENT_POS: {
			eventName: "percent-pos",
			trigger:   ipc.triggerPercentPos,
		},
		protos.Message_MESSAGE_MPV_TIME_POS: {
			eventName: "time-pos",
			trigger:   ipc.triggerTimePos,
		},
		protos.Message_MESSAGE_MPV_TIME_REMAINING: {
			eventName: "time-remaining",
			trigger:   ipc.triggerTimeRemaining,
		},
		protos.Message_MESSAGE_MPV_EOF_REACHED: {
			eventName: "eof-reached",
			trigger:   ipc.triggerStartsEnds,
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
		log.Println("connection ipc/mpv has been established")
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
			dispatcher, ok := m.listeners[protos.Message(v.ID)]
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

func (m *IPC) Load(media *protos.Media) error {
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

func (m *IPC) SeekAbsolute(percent float64) error {
	cmd := cmdSeekAbsolute
	_, err := m.connection.Call(cmd.String(), percent, "absolute-percent")
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
	return nil
}

func (m *IPC) Skip() {
	if m.lastMedia == nil {
		return
	}
	m.control.OnSongEnded(m.lastMedia, true)
}
