package broker

import (
	"encoding/json"
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

func PublishMediaUpdateDb(nc *nats.Conn, media *pb.Media) error {
	return PublishBroker(nc, pb.Message_MESSAGE_MEDIA_UPDATE, &pb.MediaInfoInput{Media: media})
}

func PublishMediaTagUpdate(nc *nats.Conn, media *pb.Media) error {
	return PublishBroker(nc, pb.Message_MESSAGE_TAG_UPDATE, media)
}

func PublishMediaSHAUpdate(nc *nats.Conn, input *pb.MediaUpdateSHAInput) error {
	return PublishBroker(nc, pb.Message_MESSAGE_MEDIA_UPDATE_SHA, input)
}

func PublishMediaSHAScan(nc *nats.Conn, input *pb.SHAScanInput) error {
	return PublishBroker(nc, pb.Message_MESSAGE_SHA_SCAN, input)
}

func PublishBrokerJSON(nc *nats.Conn, subject pb.Message, input proto.Message) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return nc.Publish(strconv.Itoa(int(subject)), data)
}

func PublishMediaArtworkUpdate(nc *nats.Conn, media *pb.Media) error {
	return PublishBroker(nc, pb.Message_MESSAGE_MEDIA_UPDATE_ARTWORK, &pb.ArtworkUpdateInput{Media: media})
}

func PublishMediaEventUpsert(nc *nats.Conn, input *pb.MediaEventInput) error {
	return PublishBroker(nc, pb.Message_MESSAGE_UPSERT_MEDIA_EVENT, input)
}
