package helpers

import (
	"errors"
	"time"

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
	ok, err := RetryFunc("connecting to NATS", retries, delay, fn)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("[ERROR] could not connect to NATS")
	}
	return conn, nil
}
