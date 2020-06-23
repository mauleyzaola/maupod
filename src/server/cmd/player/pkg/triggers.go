package pkg

// TODO: create an object which has some "intelligence" like knowing which is the current track played
// throttling to avoid sending too often messages to listeners
// if track reaches certain % consider it has been played
// consider knowing when is a skip and when it isn't

func triggerTimePos(v interface{}) {
	//log.Println("triggerTimePos: ", v)
	// cast to float64
}

func triggerFilename(v interface{}) {
	//log.Println("triggerFilename: ", v)
	// cast to string
}

func triggerStreamPos(v interface{}) {
	//log.Println("triggerStreamPos: ", v)
	// cast to float64
}

func triggerStreamEnd(v interface{}) {
	//log.Println("triggerStreamEnd: ", v)
	// cast to float64
}

func triggerPercentPos(v interface{}) {
	//log.Println("triggerPercentPos: ", v)
	// cast to float64
}

func triggerTimeRemaining(v interface{}) {
	//log.Println("triggerTimeRemaining: ", v)
	// cast to float64
}
