package pkg

import (
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/protos"
)

const (
	timePosThresholdSecs = 0.5
	percentToBeCompleted = 95
	percentToBeSkipped   = 5
)

// PlayerControl is a bridge between the mpv events and maupod events
type PlayerControl struct {
	publishFn       broker.PublisherFunc
	publishFnJSON   broker.PublisherFuncJSON
	requestFn       broker.RequestFunc
	m               *protos.Media
	lastTimePos     float64
	lastPercentPos  float64
	lastIsCompleted bool // true when based in the percent position, track is assumed to be complete
}

func NewPlayerControl(publishFn broker.PublisherFunc, publishFnJSON broker.PublisherFuncJSON, requestFn broker.RequestFunc) *PlayerControl {
	p := &PlayerControl{
		publishFn:     publishFn,
		publishFnJSON: publishFnJSON,
		requestFn:     requestFn,
	}
	return p
}

func (p *PlayerControl) OnSongEnded(m *protos.Media, isSkip bool) {
	// check if track has completely played
	if !isSkip {
		// send message to notify track has completely played
		_ = p.publishFn(protos.Message_MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE, &protos.TrackPlayedInput{
			Media:     m,
			Timestamp: helpers.TimeToTs(helpers.Now()),
		})
		p.lastIsCompleted = true

	} else {
		//  send message to notify track has been skipped
		_ = p.publishFn(protos.Message_MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE, &protos.TrackPlayedInput{
			Media:     m,
			Timestamp: helpers.TimeToTs(helpers.Now()),
		})
	}
	var output protos.QueueOutput
	if err := p.requestFn(protos.Message_MESSAGE_QUEUE_LIST, &protos.QueueInput{}, &output); err != nil {
		log.Println(err)
		return
	}

	// check queue is not empty
	if len(output.Rows) == 0 {
		log.Println("reached end of queue")
		return
	}

	// play next song in the queue
	var input = output.Rows[0]
	var ipcInput = protos.IPCInput{
		Media:   input.Media,
		Command: protos.Message_IPC_PLAY,
	}
	if err := p.publishFn(protos.Message_MESSAGE_IPC, &ipcInput); err != nil {
		log.Println(err)
		return
	}

	// remove the first element from the queue
	log.Println("[DEBUG] remove from queue: ", input.Media.Track)
	var queueInput = &protos.QueueInput{
		Media: input.Media,
		Index: 0,
	}
	if err := p.publishFn(protos.Message_MESSAGE_QUEUE_REMOVE, queueInput); err != nil {
		log.Println(err)
		return
	}
}

func (p *PlayerControl) OnSongStarted(m *protos.Media) {
	// TODO: how can this happen? it does :(
	if m == nil {
		return
	}

	// initialize values
	p.lastPercentPos = 0
	p.m = m
	p.lastTimePos = 0
	p.lastIsCompleted = false
	log.Printf("OnSongStarted id: %v track: %v\n", p.m.Id, p.m.Track)
	input := &protos.TrackStartedInput{
		Media:     p.m,
		Timestamp: helpers.TimeToTs(helpers.Now()),
	}
	_ = p.publishFn(protos.Message_MESSAGE_EVENT_ON_TRACK_STARTED, input)

	// send message to the UI through websockets
	_ = p.publishFnJSON(protos.Message_MESSAGE_SOCKET_PLAY_TRACK, &protos.PlayTrackInput{Media: m})
}

func (p *PlayerControl) onTimePosChanged(v float64) {
	// evaluate how often we want this event to be triggered
	if v == 0 {
		return
	}
	p.lastTimePos = v
	diff := v - p.lastTimePos
	if diff >= timePosThresholdSecs {
		p.OnTimePosChanged(v)
	}
}

func (p *PlayerControl) OnTimePosChanged(v float64) {
	//log.Println("OnTimePosChanged: ", v)
}

func (p *PlayerControl) onPercentPosChanged(media *protos.Media, v float64) {
	if media == nil {
		return
	}
	p.lastPercentPos = v
	p.OnPercentPosChanged(media, v)
}

func (p *PlayerControl) OnPercentPosChanged(media *protos.Media, v float64) {
	// location should be relative
	tmpMedia := *media
	tmpMedia.Location = paths.LocationPath(tmpMedia.Location)
	secondsTotal, err := rules.MediaTotalSeconds(media)
	if err != nil {
		log.Println(err)
		return
	}
	// we need to send json here, so easier to deal for node
	if err := p.publishFnJSON(protos.Message_MESSAGE_SOCKET_TRACK_POSITION_PERCENT, &protos.TrackPositionInput{
		Media:        &tmpMedia,
		Percent:      v,
		Seconds:      p.lastTimePos,
		SecondsTotal: secondsTotal.Seconds(),
	}); err != nil {
		log.Println(err)
	}
}
