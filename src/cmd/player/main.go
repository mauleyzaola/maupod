package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/broker"

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

	// we're outside the docker network, need to hit localhost
	config.NatsUrl = nats.DefaultURL

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