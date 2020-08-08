package broker

import (
	"log"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

func MicroServiceRespond(msg *nats.Msg, name string) error {
	var err error
	var input pb.MicroServiceDiscoveryInput
	var output pb.MicroServiceDiscoveryOutput
	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return err
	}

	output.Ok = true
	output.Name = name

	data, err := helpers.ProtoMarshal(&output)
	if err != nil {
		log.Println(err)
		return err
	}
	if err = msg.Respond(data); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func MicroServicePing(nc *nats.Conn, name pb.Message, timeout time.Duration) (*pb.MicroServiceDiscoveryOutput, error) {
	var input = pb.MicroServiceDiscoveryInput{}
	var output pb.MicroServiceDiscoveryOutput
	if err := DoRequest(nc, name, &input, &output, timeout); err != nil {
		return nil, err
	}
	return &output, nil
}
