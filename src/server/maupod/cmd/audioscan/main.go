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
	config.MediaStores = rule.ConfigurationFileSystemStores(config)
	if len(config.MediaStores) == 0 {
		return errors.New("could not find any media store in configuration or environment")
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
	hnd = NewMsgHandler(config, db, logger, nc)
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
