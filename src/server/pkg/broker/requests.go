package broker

import (
	"errors"
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
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

func RequestMediaInfoScan(nc *nats.Conn, logger types.Logger, filename string, timeout time.Duration) (*pb.Media, error) {
	input := &pb.MediaInfoInput{FileName: filename}
	output, err := mediaInfoRequest(nc, input, timeout)
	if err != nil {
		// TODO: send the files with errors to another listener and store in db
		logger.Error(err)
		return nil, err
	}
	if output.Response == nil {
		return nil, errors.New("missing response")
	}
	if !output.Response.Ok {
		return nil, errors.New(output.Response.Error)
	}
	if output.Media == nil {
		return nil, errors.New("mediainfo returned a nil object")
	}
	output.Media.ModifiedDate = output.LastModifiedDate

	return output.Media, nil
}
