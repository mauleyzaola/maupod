package pkg

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
	//log.Println("triggerStreamPos: ", v)
	// cast to float64
}

func (m *IPC) triggerStreamEnd(v interface{}) {
	//log.Println("triggerStreamEnd: ", v)
	// cast to float64
}

func (m *IPC) triggerPercentPos(v interface{}) {
	//log.Println("triggerPercentPos: ", v)
	// cast to float64
}

func (m *IPC) triggerTimeRemaining(v interface{}) {
	//log.Println("triggerTimeRemaining: ", v)
	// cast to float64
}
