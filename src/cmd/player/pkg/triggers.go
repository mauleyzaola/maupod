package pkg

import (
	"log"
	"time"
)

// TODO: create an object which has some "intelligence" like knowing which is the current track played
// throttling to avoid sending too often messages to listeners
// if track reaches certain % consider it has been played
// consider knowing when is a skip and when it isn't

func (m *IPC) triggerTimePos(v interface{}) {
	val, ok := v.(float64)
	if !ok {
		return
	}
	m.control.onTimePosChanged(val)
}

func (m *IPC) triggerStreamPos(v interface{}) {
	log.Println(v)
}

func (m *IPC) triggerStreamEnd(v interface{}) {}

func (m *IPC) triggerPercentPos(v interface{}) {
	val, ok := v.(float64)
	if !ok {
		return
	}
	if m.playedStateFn != nil {
		m.playedStateFn(m.lastMedia, val)
	}
	m.control.onPercentPosChanged(m.lastMedia, val)
}

func (m *IPC) triggerTimeRemaining(v interface{}) {
	//log.Println("triggerTimeRemaining: ", v)
	// cast to float64
}

func (m *IPC) triggerStartsEnds(v interface{}) {
	if m.connection == nil {
		log.Println("[WARNING] m.connection not available yet")
		return
	}
	val, ok := v.(bool)
	if !ok {
		return
	}

	// workaround triggering a second time the play track event
	const threshold = time.Millisecond * 100
	var now = time.Now()
	var diff = now.Sub(m.lastStartTrackEvent)
	m.lastStartTrackEvent = now
	if diff < threshold {
		return
	}

	if val {
		if m.stoppedState != nil {
			m.stoppedState()
		}
		m.control.OnSongEnded(m.lastMedia, false)
	} else {
		m.control.OnSongStarted(m.lastMedia)
	}

}
