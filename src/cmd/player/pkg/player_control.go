package pkg

import (
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/rules"

	"github.com/mauleyzaola/maupod/src/pkg/paths"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

const (
	timePosThresholdSecs = 0.5
	percentToBeCompleted = 50
	percentToBeSkipped   = 5
)

// PlayerControl is a bridge between the mpv events and maupod events
type PlayerControl struct {
	publishFn       broker.PublisherFunc
	publishFnJSON   broker.PublisherFuncJSON
	requestFn       broker.RequestFunc
	m               *pb.Media
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

func (p *PlayerControl) OnSongEnded(m *pb.Media) {
	var output pb.QueueOutput
	if err := p.requestFn(pb.Message_MESSAGE_QUEUE_LIST, &pb.QueueInput{}, &output); err != nil {
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
	var ipcInput = pb.IPCInput{
		Media:   input.Media,
		Command: pb.Message_IPC_PLAY,
	}
	if err := p.publishFn(pb.Message_MESSAGE_IPC, &ipcInput); err != nil {
		log.Println(err)
		return
	}

	// remove the first element from the queue
	log.Println("[DEBUG] remove from queue: ", input.Media.Track)
	var queueInput = &pb.QueueInput{
		Media: input.Media,
		Index: 0,
	}
	if err := p.publishFn(pb.Message_MESSAGE_QUEUE_REMOVE, queueInput); err != nil {
		log.Println(err)
		return
	}
}

func (p *PlayerControl) OnSongStarted(m *pb.Media) {
	// TODO: how can this happen? it does :(
	if m == nil {
		return
	}
	// read state
	var isNewTrack = p.m == nil || p.m.Id != m.Id
	var lastPercentPos = p.lastPercentPos

	// initialize values
	p.lastPercentPos = 0
	p.m = m
	p.lastTimePos = 0
	p.lastIsCompleted = false
	log.Printf("OnSongStarted id: %v track: %v\n", p.m.Id, p.m.Track)
	input := &pb.TrackStartedInput{
		Media:     p.m,
		Timestamp: helpers.TimeToTs(helpers.Now()),
	}
	_ = p.publishFn(pb.Message_MESSAGE_EVENT_ON_TRACK_STARTED, input)

	if isNewTrack {
		if lastPercentPos >= percentToBeSkipped && lastPercentPos < percentToBeCompleted {
			input := &pb.TrackSkippedInput{
				Media:     m,
				Timestamp: helpers.TimeToTs(helpers.Now()),
			}
			_ = p.publishFn(pb.Message_MESSAGE_EVENT_ON_TRACK_SKIP_COUNT_INCREASE, input)
		}
	}
	// send message to the UI through websockets
	_=p.publishFnJSON(pb.Message_MESSAGE_SOCKET_PLAY_TRACK, &pb.PlayTrackInput{Media: m})
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

func (p *PlayerControl) onPercentPosChanged(media *pb.Media, v float64) {
	if media == nil {
		return
	}
	p.lastPercentPos = v
	p.OnPercentPosChanged(media, v)
}

func (p *PlayerControl) OnPercentPosChanged(media *pb.Media, v float64) {
	// check percente position to know if track has completed playing
	if !p.lastIsCompleted {
		if v >= percentToBeCompleted {
			_ = p.publishFn(pb.Message_MESSAGE_EVENT_ON_TRACK_PLAY_COUNT_INCREASE, &pb.TrackPlayedInput{
				Media:     media,
				Timestamp: helpers.TimeToTs(helpers.Now()),
			})
			p.lastIsCompleted = true
		}
	}

	// location should be relative
	tmpMedia := *media
	tmpMedia.Location = paths.LocationPath(tmpMedia.Location)
	secondsTotal, err := rules.MediaTotalSeconds(media)
	if err != nil {
		log.Println(err)
		return
	}
	// we need to send json here, so easier to deal for node
	if err := p.publishFnJSON(pb.Message_MESSAGE_SOCKET_TRACK_POSITION_PERCENT, &pb.TrackPositionInput{
		Media:        &tmpMedia,
		Percent:      float32(v),
		Seconds:      float32(p.lastTimePos),
		SecondsTotal: float32(secondsTotal.Seconds()),
	}); err != nil {
		log.Println(err)
	}
}
