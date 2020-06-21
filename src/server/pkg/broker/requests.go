package broker

import (
	"errors"
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func mediaInfoRequest(nc *nats.Conn, input *pb.MediaInfoInput, timeout time.Duration) (*pb.MediaInfoOutput, error) {
	var output pb.MediaInfoOutput
	data, err := helpers.ProtoMarshal(input)
	if err != nil {
		return nil, err
	}
	msg, err := nc.Request(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_INFO)), data, timeout)
	if err != nil {
		return nil, err
	}

	if err = helpers.ProtoUnmarshal(msg.Data, &output); err != nil {
		return nil, err
	}

	return &output, nil
}

func doRequest(nc *nats.Conn, subject int, input, output proto.Message, timeout time.Duration) error {
	data, err := helpers.ProtoMarshal(input)
	if err != nil {
		return err
	}
	msg, err := nc.Request(strconv.Itoa(subject), data, timeout)
	if err != nil {
		return err
	}

	if err = helpers.ProtoUnmarshal(msg.Data, output); err != nil {
		return err
	}

	return nil
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

func RequestIPCCommand(nc *nats.Conn, input *pb.IPCInput, timeout time.Duration) (*pb.IPCOutput, error) {
	var output = &pb.IPCOutput{}
	if err := doRequest(nc, int(pb.Message_MESSAGE_IPC), input, output, timeout); err != nil {
		return nil, err
	}
	return output, nil
}
