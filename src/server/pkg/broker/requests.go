package broker

import (
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func mediaInfoRequest(nc *nats.Conn, input *pb.MediaInfoInput, timeout time.Duration) (*pb.MediaInfoOutput, error) {
	var output pb.MediaInfoOutput
	data, err := proto.Marshal(input)
	if err != nil {
		return nil, err
	}
	msg, err := nc.Request(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_INFO)), data, timeout)
	if err != nil {
		return nil, err
	}

	if err = proto.Unmarshal(msg.Data, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

func PublishMediaInfoDelete(nc *nats.Conn, input *pb.MediaInfoInput) error {
	data, err := proto.Marshal(input)
	if err != nil {
		return err
	}
	return nc.Publish(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_DELETE)), data)
}

// TODO: implement the caller
func PublishMediaSHAUpdate(nc *nats.Conn, input *pb.MediaInfoInput) error {
	fileData, err := proto.Marshal(input)
	if err != nil {
		return err
	}
	return PublishMessage(nc, pb.Message_MESSAGE_AUDIO_SHA, fileData)
}
