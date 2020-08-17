package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
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

	// don't use common NATS port
	config.NatsUrl = "nats://127.0.0.1:4244"

	nc, err := broker.ConnectNATS(config.NatsUrl, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	log.Println("successfully connected to NATS")

	var hnd types.Broker
	hnd = NewMsgHandler(nc)
	if err = hnd.Register(); err != nil {
		return err
	}

	// check for dependent micro services to be up
	timeout := time.Second * time.Duration(config.Delay)
	var dependentMicroServices = []pb.Message{
		pb.Message_MESSAGE_MICRO_SERVICE_RESTAPI,
		pb.Message_MESSAGE_MICRO_SERVICE_MEDIAINFO,
		pb.Message_MESSAGE_MICRO_SERVICE_SOCKET,
	}
	for _, v := range dependentMicroServices {
		helpers.RetryFunc(v.String(), int(config.Retries), timeout, func(retryCount int) bool {
			if _, err := broker.MicroServicePing(nc, pb.Message_MESSAGE_MICRO_SERVICE_RESTAPI, timeout); err != nil {
				return false
			}
			return true
		})
	}

	// TODO: configure this feature somewhere else
	if err := autoPlayQueue(nc, config); err != nil {
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

// TODO: implement retry logic in case nats listener is not available yet
func autoPlayQueue(nc *nats.Conn, config *pb.Configuration) error {
	var timeout = rules.Timeout(config)
	output, err := broker.RequestQueueList(nc, &pb.QueueInput{}, timeout)
	if err != nil {
		return err
	}
	if len(output.Rows) == 0 {
		return nil
	}
	next := output.Rows[0]
	// send a nats message
	var input = &pb.IPCInput{
		Media:   next.Media,
		Value:   "",
		Command: pb.Message_IPC_PLAY,
	}
	if err = broker.RequestIPCCommand(nc, input, timeout); err != nil {
		return err
	}
	if _, err = broker.RequestQueueRemove(nc, &pb.QueueInput{
		Media: next.Media,
		Index: 0,
	}, timeout); err != nil {
		return err
	}
	return nil
}
