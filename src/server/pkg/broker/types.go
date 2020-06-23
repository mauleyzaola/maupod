package broker

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

// generic function definition to mock NATS publish behavior
type PublisherFunc func(nc *nats.Conn, subject pb.Message, data []byte) error
