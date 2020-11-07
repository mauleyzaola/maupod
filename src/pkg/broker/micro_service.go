package broker

import (
	"log"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/nats-io/nats.go"
)

func MicroServiceRespond(msg *nats.Msg, name string) error {
	var err error
	var input protos.MicroServiceDiscoveryInput
	var output protos.MicroServiceDiscoveryOutput
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

func MicroServicePing(nc *nats.Conn, name protos.Message, timeout time.Duration) (*protos.MicroServiceDiscoveryOutput, error) {
	var input = protos.MicroServiceDiscoveryInput{}
	var output protos.MicroServiceDiscoveryOutput
	if err := DoRequest(nc, name, &input, &output, timeout); err != nil {
		return nil, err
	}
	return &output, nil
}
