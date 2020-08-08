package broker

import (
	"errors"
	"log"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/pb"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/nats-io/nats.go"
)

func ConnectNATS(natsURL string, retries int, delay time.Duration) (*nats.Conn, error) {
	if natsURL == "" {
		return nil, errors.New("cannot resolve variable: NATS_URL")
	}
	var conn *nats.Conn
	var err error
	fn := func(retry int) bool {
		if conn, err = nats.Connect(natsURL); err != nil {
			return false
		}
		return conn != nil
	}
	ok, err := helpers.RetryFunc("connecting to NATS", retries, delay, fn)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("[ERROR] could not connect to NATS")
	}
	return conn, nil
}

func RestAPIPing(nc *nats.Conn, retries int, delay, timeout time.Duration) error {
	var ok bool
	fn := func(retry int) bool {
		if _, err := MicroServicePing(nc, pb.Message_MESSAGE_MICRO_SERVICE_RESTAPI, timeout); err != nil {
			return false
		}
		ok = true
		return ok
	}
	if _, err := helpers.RetryFunc("ping RestAPI", retries, delay, fn); err != nil {
		return err
	}
	if !ok {
		return errors.New("could not ping RestAPI")
	}
	log.Println("[INFO] successfully ping to RestAPI from: ", helpers.AppName())
	return nil
}
