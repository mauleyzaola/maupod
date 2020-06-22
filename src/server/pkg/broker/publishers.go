package broker

import (
	"strconv"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func PublishMessage(nc *nats.Conn, subject pb.Message, data []byte) error {
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
	fileData, err := helpers.ProtoMarshal(input)
	if err != nil {
		return err
	}
	return PublishMessage(nc, pb.Message_MESSAGE_AUDIO_SHA, fileData)
}

func PublishMediaUpdateDb(nc *nats.Conn, media *pb.Media) error {
	data, err := helpers.ProtoMarshal(media)
	if err != nil {
		return err
	}
	return PublishMessage(nc, pb.Message_MESSAGE_MEDIA_UPDATE, data)
}

func PublishMediaTagUpdate(nc *nats.Conn, media *pb.Media) error {
	data, err := helpers.ProtoMarshal(media)
	if err != nil {
		return err
	}
	return PublishMessage(nc, pb.Message_MESSAGE_TAG_UPDATE, data)
}
