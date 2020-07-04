package helpers

import (
	"errors"
	"time"

	"github.com/golang/glog"
	"gopkg.in/redis.v4"
)

func ConnectRedis(host string, retry int, delay time.Duration) (*redis.Client, error) {
	const port = ":6379"
	client := redis.NewClient(&redis.Options{
		Addr:     host + port,
		Password: "",
		DB:       0,
	})
	glog.V(1).Infoln("[INFO] trying to ping the redis server from container")

	ok, err := RetryFunc("connecting to redis", retry, delay, func(int) bool {
		if err := client.Ping(); err != nil {
			return false
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("could not connect to redis server")
	}
	return client, nil
}
