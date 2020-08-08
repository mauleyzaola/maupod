package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/pb"

	"github.com/mauleyzaola/maupod/src/pkg/broker"

	_ "github.com/lib/pq"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/mauleyzaola/maupod/src/pkg/simplelog"
	"github.com/mauleyzaola/maupod/src/pkg/types"
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

	config, err := rules.ConfigurationParse()
	if err != nil {
		return err
	}

	if err = rules.ConfigurationValidate(config); err != nil {
		return err
	}

	// create directory if not exists
	if val := os.Getenv("MEDIA_STORE"); val != "" {
		config.MediaStores = []*pb.FileStore{
			{
				Type:     pb.FileStore_FILE_SYSTEM,
				Name:     "media-store",
				Location: val,
			},
		}
	}

	if len(config.MediaStores) == 0 {
		return errors.New("could not find any media store in configuration or environment")
	}

	nc, err := broker.ConnectNATS(config.NatsUrl, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	logger.Info("successfully connected to NATS")

	delay := time.Second * time.Duration(config.Delay)
	if err = broker.RestAPIPing(nc, int(config.Retries), delay, delay); err != nil {
		return err
	}

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
