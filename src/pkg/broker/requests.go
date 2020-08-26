package broker

import (
	"errors"
	"strconv"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func DoRequest(nc *nats.Conn, subject pb.Message, input, output proto.Message, timeout time.Duration) error {
	data, err := helpers.ProtoMarshal(input)
	if err != nil {
		return err
	}
	msg, err := nc.Request(strconv.Itoa(int(subject)), data, timeout)
	if err != nil {
		return err
	}
	if output != nil {
		if err = helpers.ProtoUnmarshal(msg.Data, output); err != nil {
			return err
		}
	}
	return nil
}

func doPublish(nc *nats.Conn, subject pb.Message, input proto.Message) error {
	data, err := helpers.ProtoMarshal(input)
	if err != nil {
		return err
	}
	return nc.Publish(strconv.Itoa(int(subject)), data)
}

func RequestMediaInfoScan(nc *nats.Conn, filename string, timeout time.Duration) (*pb.MediaInfoOutput, error) {
	var output pb.MediaInfoOutput
	input := &pb.MediaInfoInput{FileName: filename}
	if err := DoRequest(nc, pb.Message_MESSAGE_MEDIA_INFO, input, &output, timeout); err != nil {
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

	return &output, nil
}

func RequestMediaInfoScanFromDB(nc *nats.Conn, input *pb.MediaInfoInput, timeout time.Duration) (*pb.MediaInfosOutput, error) {
	var output pb.MediaInfosOutput
	if err := DoRequest(nc, pb.Message_MESSAGE_MEDIA_DB_SELECT, input, &output, timeout); err != nil {
		return nil, err
	}
	if output.Response == nil {
		return nil, errors.New("missing response")
	}
	if !output.Response.Ok {
		return nil, errors.New(output.Response.Error)
	}
	return &output, nil
}

func RequestIPCCommand(nc *nats.Conn, input *pb.IPCInput, timeout time.Duration) error {
	return doPublish(nc, pb.Message_MESSAGE_IPC, input)
}

func RequestQueueAdd(nc *nats.Conn, input *pb.QueueInput, timeout time.Duration) (*pb.QueueOutput, error) {
	var output pb.QueueOutput
	if err := DoRequest(nc, pb.Message_MESSAGE_QUEUE_ADD, input, &output, timeout); err != nil {
		return nil, err
	}
	if output.Error != "" {
		return nil, errors.New(output.Error)
	}
	return &output, nil
}

func RequestQueueRemove(nc *nats.Conn, input *pb.QueueInput, timeout time.Duration) (*pb.QueueOutput, error) {
	var output pb.QueueOutput
	if err := DoRequest(nc, pb.Message_MESSAGE_QUEUE_REMOVE, input, &output, timeout); err != nil {
		return nil, err
	}
	if output.Error != "" {
		return nil, errors.New(output.Error)
	}
	return &output, nil
}

func RequestQueueList(nc *nats.Conn, input *pb.QueueInput, timeout time.Duration) (*pb.QueueOutput, error) {
	var output pb.QueueOutput
	if err := DoRequest(nc, pb.Message_MESSAGE_QUEUE_LIST, input, &output, timeout); err != nil {
		return nil, err
	}
	if output.Error != "" {
		return nil, errors.New(output.Error)
	}
	return &output, nil
}

func RequestFileBrowserDirectory(nc *nats.Conn, input *pb.DirectoryReadInput, timeout time.Duration) (*pb.DirectoryReadOutput, error) {
	var output pb.DirectoryReadOutput
	if err := DoRequest(nc, pb.Message_MESSAGE_DIRECTORY_READ, input, &output, timeout); err != nil {
		return nil, err
	}
	return &output, nil
}
