package pkg

import "log"

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
	m.control.onPercentPosChanged(m.lastMedia, val)
}

func (m *IPC) triggerTimeRemaining(v interface{}) {
	//log.Println("triggerTimeRemaining: ", v)
	// cast to float64
}

func (m *IPC) triggerStartsEnds(v interface{}) {
	val, ok := v.(bool)
	if !ok {
		return
	}
	if val {
		m.control.OnSongEnded(m.lastMedia)
	} else {
		m.control.OnSongStarted(m.lastMedia)
	}

}
