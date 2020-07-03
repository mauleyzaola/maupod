package broker

import (
	"errors"
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func doRequest(nc *nats.Conn, subject pb.Message, input, output proto.Message, timeout time.Duration) error {
	data, err := helpers.ProtoMarshal(input)
	if err != nil {
		return err
	}
	msg, err := nc.Request(strconv.Itoa(int(subject)), data, timeout)
	if err != nil {
		return err
	}
	if err = helpers.ProtoUnmarshal(msg.Data, output); err != nil {
		return err
	}
	return nil
}

func RequestMediaInfoScan(nc *nats.Conn, logger types.Logger, filename string, timeout time.Duration) (*pb.Media, error) {
	var output pb.MediaInfoOutput
	input := &pb.MediaInfoInput{FileName: filename}
	if err := doRequest(nc, pb.Message_MESSAGE_MEDIA_INFO, input, &output, timeout); err != nil {
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

func RequestIPCCommand(nc *nats.Conn, input *pb.IPCInput, timeout time.Duration) (*pb.IPCOutput, error) {
	var output = &pb.IPCOutput{}
	if err := doRequest(nc, pb.Message_MESSAGE_IPC, input, output, timeout); err != nil {
		return nil, err
	}
	return output, nil
}

func RequestRestAPIReady(nc *nats.Conn, timeout time.Duration) error {
	_, err := nc.Request(strconv.Itoa(int(pb.Message_MESSAGE_REST_API_READY)), nil, timeout)
	return err
}