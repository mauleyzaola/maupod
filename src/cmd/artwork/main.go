package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	_ "github.com/lib/pq"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/pkg/types"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
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

	// create directory if not exists
	if config.ArtworkStore != nil {
		// TODO: probably other store types should not need to create the directory, use an interface instead
		if err = os.MkdirAll(config.ArtworkStore.Location, os.ModePerm); err != nil {
			return err
		}
	} else {
		return errors.New("could not find any image store in configuration")
	}

	nc, err := broker.ConnectNATS(config.NatsUrl, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	log.Println("successfully connected to NATS")

	delay := time.Second * time.Duration(config.Delay)
	if err = broker.RestAPIPing(nc, int(config.Retries), delay, delay); err != nil {
		return err
	}

	db, err := dbdata.ConnectPostgres(config.DbConn, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}

	var hnd types.Broker
	hnd = NewMsgHandler(db, config, nc)
	if err = hnd.Register(); err != nil {
		return err
	}

	// handle interruptions and cleanup resources
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		cleanup := func() {
			log.Println("received an interrupt signal, cleaning resources...")
			hnd.Close()

			log.Println("completed cleaning up resources")
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
