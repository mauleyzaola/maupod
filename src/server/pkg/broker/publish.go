package broker

import "github.com/nats-io/nats.go"

func PublishMessage(nc *nats.Conn, subject string, data []byte) error {
	return nc.Publish(subject, data)
}
