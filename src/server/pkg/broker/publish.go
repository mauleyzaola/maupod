package broker

import (
	"strconv"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func PublishMessage(nc *nats.Conn, subject pb.Message, data []byte) error {
	return nc.Publish(strconv.Itoa(int(subject)), data)
}
