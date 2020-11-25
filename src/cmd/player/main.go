package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/protos"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	helpers.AppInit()
}

func run() error {
	config, err := rules.ConfigurationParse()
	if err != nil {
		return err
	}

	if err = rules.ConfigurationValidate(config); err != nil {
		return err
	}

	timeout := time.Second * time.Duration(config.Delay)

	// connect to redis
	var rc *redis.Client
	// TODO: support ip:port address connection
	host, port := "localhost", "14100"
	if _, err = helpers.RetryFunc(fmt.Sprintf("connecting to redis %s:%s\n", host, port), int(config.Retries), timeout, func(retryCount int) bool {
		rc, err = dbdata.ConnectRedis(host, port)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}); err != nil {
		log.Println("could not connect to redis")
		return err
	}

	// don't use common NATS port
	// TODO: support ip address
	config.NatsUrl = "nats://127.0.0.1:4244"

	nc, err := broker.ConnectNATS(config.NatsUrl, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	log.Println("successfully connected to NATS")

	hnd := NewMsgHandler(nc, rc, timeout)
	if err = hnd.Register(); err != nil {
		return err
	}

	// check for dependent micro services to be up
	var dependentMicroServices = map[protos.Message]bool{
		protos.Message_MESSAGE_MICRO_SERVICE_RESTAPI:   false,
		protos.Message_MESSAGE_MICRO_SERVICE_MEDIAINFO: false,
	}

	for k := range dependentMicroServices {
		_, _ = helpers.RetryFunc(k.String(), int(config.Retries), timeout, func(retryCount int) bool {
			if _, err := broker.MicroServicePing(nc, k, timeout); err != nil {
				return false
			}
			dependentMicroServices[k] = true
			return true
		})
	}
	for k, v := range dependentMicroServices {
		if !v {
			err = fmt.Errorf("[ERROR] could not establish connection to micro service: %s\n", k)
			return err
		}
	}
	log.Println("[INFO] all micro services are up")

	// start handler in case we need to resume playing
	if err = hnd.Start(); err != nil {
		log.Println(err)
	}

	// handle interruptions and cleanup resources
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		cleanup := func() {
			hnd.Close()
			cleanupDone <- true
		}
	loop:
		for {
			select {
			case <-signalChan:
				break loop
			}
		}
		cleanup()
	}()

	<-cleanupDone

	return nil
}
