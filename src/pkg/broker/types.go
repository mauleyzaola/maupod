package broker

import (
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"google.golang.org/protobuf/proto"
)

// generic function definition to mock NATS publish behavior
type PublisherFunc func(subject pb.Message, payload proto.Message) error

type RequestFunc func(subject pb.Message, input, output proto.Message) error
