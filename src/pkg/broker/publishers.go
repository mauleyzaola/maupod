package broker

import (
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func PublishBroker(nc *nats.Conn, subject pb.Message, input proto.Message) error {
	data, err := helpers.ProtoMarshal(input)
	if err != nil {
		return err
	}
	return nc.Publish(strconv.Itoa(int(subject)), data)
}

func PublishMediaInfoDelete(nc *nats.Conn, input *pb.MediaInfoInput) error {
	data, err := helpers.ProtoMarshal(input)
	if err != nil {
		return err
	}
	return nc.Publish(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_DELETE)), data)
}

func PublishMediaSHAUpdate(nc *nats.Conn, input *pb.MediaInfoInput) error {
	return PublishBroker(nc, pb.Message_MESSAGE_AUDIO_SHA, input)
}

func PublishMediaUpdateDb(nc *nats.Conn, media *pb.Media) error {
	return PublishBroker(nc, pb.Message_MESSAGE_MEDIA_UPDATE, media)
}

func PublishMediaTagUpdate(nc *nats.Conn, media *pb.Media) error {
	return PublishBroker(nc, pb.Message_MESSAGE_TAG_UPDATE, media)
}
