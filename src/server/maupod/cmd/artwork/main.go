package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/lib/pq"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/mauleyzaola/maupod/src/server/pkg/simplelog"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/spf13/viper"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
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

	config, err := rule.ConfigurationParse()
	if err != nil {
		return err
	}

	if err = rule.ConfigurationValidate(config); err != nil {
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

	nc, err := helpers.ConnectNATS(config.NatsUrl, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	logger.Info("successfully connected to NATS")

	db, err := dbdata.ConnectPostgres(config.DbConn, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}

	var hnd types.Broker
	hnd = NewMsgHandler(db, config, logger, nc)
	if err = hnd.Register(); err != nil {
		return err
	}

	// handle interruptions and cleanup resources
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		cleanup := func() {
			logger.Info("received an interrupt signal, cleaning resources...")
			hnd.Close()

			logger.Info("completed cleaning up resources")
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
