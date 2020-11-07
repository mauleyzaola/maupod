package broker

import (
	"github.com/mauleyzaola/maupod/src/protos"
	"google.golang.org/protobuf/proto"
)

// generic function definition to mock NATS publish behavior
type PublisherFunc func(subject protos.Message, payload proto.Message) error
type PublisherFuncJSON func(subject protos.Message, payload interface{}) error
type RequestFunc func(subject protos.Message, input, output proto.Message) error
