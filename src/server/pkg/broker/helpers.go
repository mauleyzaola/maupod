package broker

import (
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func MediaInfoRequest(nc *nats.Conn, input *pb.MediaInfoInput, timeout time.Duration) (*pb.MediaInfoOutput, error) {
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
