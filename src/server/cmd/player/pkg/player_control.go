package pkg

import (
	"log"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/server/pkg/broker"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

const (
	timePosThresholdSecs = 0.5
	percentToBeCompleted = 80
	percentToBeSkipped   = 20
)

// PlayerControl is a bridge between the mpv events and maupod events
type PlayerControl struct {
	publishFn       broker.PublisherFunc
	m               *pb.Media
	lastTimePos     float64
	lastPercentPos  float64
	lastIsCompleted bool // true when based in the percent position, track is assumed to be complete
}

func NewPlayerControl(publishFn broker.PublisherFunc) *PlayerControl {
	p := &PlayerControl{
		publishFn: publishFn,
	}
	return p
}

func (p *PlayerControl) OnSongStarted(media *pb.Media) {
	// evaluate p.lastPercentPos to consider if this track was skipped or not
	if p.m == nil || p.m.Id != media.Id {
		if p.lastPercentPos >= percentToBeSkipped && p.lastPercentPos < percentToBeCompleted {
			input := &pb.TrackSkippedInput{
				Media:     media,
				Timestamp: helpers.TimeToTs(helpers.Now()),
			}
			_ = p.publishFn(pb.Message_MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE, input)
		}
	}

	p.lastPercentPos = 0
	p.m = media
	p.lastTimePos = 0
	p.lastIsCompleted = false
	log.Printf("OnSongStarted id: %v track: %v\n", p.m.Id, p.m.Track)
	input := &pb.TrackStartedInput{
		Media:     p.m,
		Timestamp: helpers.TimeToTs(helpers.Now()),
	}
	_ = p.publishFn(pb.Message_MESSAGE_EVENT_ON_TRACK_STARTED, input)
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
	//log.Println("OnTimePosChanged: ", v)
}

func (p *PlayerControl) onPercentPosChanged(media *pb.Media, v float64) {
	p.lastPercentPos = v
	if p.lastIsCompleted {
		return
	}
	p.OnPercentPosChanged(media, v)
}

func (p *PlayerControl) OnPercentPosChanged(media *pb.Media, v float64) {
	input := &pb.TrackPlayedInput{
		Media:     media,
		Timestamp: helpers.TimeToTs(helpers.Now()),
	}
	// if track has played halfway we consider it as been played
	if v >= percentToBeCompleted {
		_ = p.publishFn(pb.Message_MESSAGE_EVENT_ON_TRACK_PLAYED, input)
		p.lastIsCompleted = true
	}
}
