package pkg

import (
	"log"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

const (
	timePosThresholdSecs = 0.5
)

// PlayerControl is a bridge between the mpv events and maupod events
type PlayerControl struct {
	m           *pb.Media
	lastTimePos float64
}

func NewPlayerControl() *PlayerControl {
	p := &PlayerControl{}
	return p
}

func (p *PlayerControl) OnSongStarted(media *pb.Media) {
	p.m = media
	p.lastTimePos = 0
	log.Printf("OnSongStarted id: %v track: %v\n", p.m.Id, p.m.Track)
}

func (p *PlayerControl) onTimePosChanged(v float64) {
	// evaluate how often we want this event to be triggered
	if v == 0 {
		return
	}

	diff := v - p.lastTimePos
	if diff >= timePosThresholdSecs {
		p.OnTimePosChanged(v)
		p.lastTimePos = v
	}
}

func (p *PlayerControl) OnTimePosChanged(v float64) {
	log.Println("OnTimePosChanged: ", v)
}
