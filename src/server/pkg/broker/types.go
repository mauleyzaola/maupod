package broker

import (
	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

// generic function definition to mock NATS publish behavior
type PublisherFunc func(subject pb.Message, payload proto.Message) error
