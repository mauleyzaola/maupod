package broker

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"google.golang.org/protobuf/proto"
)

// generic function definition to mock NATS publish behavior
type PublisherFunc func(subject pb.Message, payload proto.Message) error
