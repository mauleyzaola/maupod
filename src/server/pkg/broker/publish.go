package broker

import (
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func PublishMessage(nc *nats.Conn, subject pb.Message, data []byte) error {
	return nc.Publish(strconv.Itoa(int(subject)), data)
}

func DoNATSRequest(nc *nats.Conn, subject pb.Message, timeout time.Duration, input *proto.Message) ([]byte, error) {
	var err error
	var msgData []byte
	if input != nil {
		if msgData, err = proto.Marshal(*input); err != nil {
			return nil, err
		}
	}
	msg, err := nc.Request(strconv.Itoa(int(subject)), msgData, timeout)
	if err != nil {
		return nil, err
	}
	return msg.Data, nil
}
