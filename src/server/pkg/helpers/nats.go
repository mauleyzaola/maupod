package helpers

import (
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

func ConnectNATS() (*nats.Conn, error) {
	natsURL := viper.GetString("NATS_URL")
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
	retries := viper.GetInt("MAX_RETRIES")
	delay := viper.GetDuration("DELAY")
	ok, err := RetryFunc("connecting to NATS", retries, delay, fn)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("[ERROR] could not connect to NATS")
	}
	return conn, nil
}
