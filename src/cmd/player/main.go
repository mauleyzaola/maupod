package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/pkg/simplelog"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".maupod")

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()
}

func run() error {
	var logger types.Logger
	logger = &simplelog.Log{}
	logger.Init()

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
	logger.Info("successfully connected to NATS")

	var hnd types.Broker
	hnd = NewMsgHandler(logger, nc)
	if err = hnd.Register(); err != nil {
		return err
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

func autoPlayQueue(nc *nats.Conn, config *pb.Configuration) error {
	var timeout = rules.Timeout(config)
	output, err := broker.RequestQueueList(nc, &pb.QueueInput{}, timeout)
	if err != nil {
		return err
	}
	if len(output.Rows) == 0 {
		return nil
	}
	nextMedia := output.Rows[0]
	// send a nats message
	var input = &pb.IPCInput{
		Media:   nextMedia,
		Value:   "",
		Command: pb.Message_IPC_PLAY,
	}
	if err = broker.RequestIPCCommand(nc, input, timeout); err != nil {
		return err
	}
	if _, err = broker.RequestQueueRemove(nc, &pb.QueueInput{
		Media: nextMedia,
		Index: 0,
	}, timeout); err != nil {
		return err
	}
	return nil
}
