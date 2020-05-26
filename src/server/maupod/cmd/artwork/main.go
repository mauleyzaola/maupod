package main

import (
	"errors"
	"log"
	"os"
	"os/signal"

	"github.com/golang/glog"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"

	"github.com/mauleyzaola/maupod/src/server/pkg/simplelog"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
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
	config, err := rule.ConfigurationParse()
	if err != nil {
		return err
	}

	if err = rule.ConfigurationValidate(config); err != nil {
		return err
	}

	// create directory if not exists
	imageStore := rule.ConfigurationFirstImageStore(config)
	if imageStore != nil {
		if err = os.MkdirAll(imageStore.Location, os.ModePerm); err != nil {
			return err
		}
	} else {
		return errors.New("could not find any image store in configuration")
	}

	nc, err := helpers.ConnectNATS()
	if err != nil {
		return err
	}

	var logger types.Logger
	logger = &simplelog.Log{}
	logger.Init()

	db, err := data.DbBootstrap(config)
	if err != nil {
		return err
	}

	var hnd types.Broker
	hnd = NewMsgHandler(db, imageStore, logger, nc)
	if err = hnd.Register(); err != nil {
		return err
	}

	// handle interruptions and cleanup resources
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		cleanup := func() {
			glog.V(1).Infof("received an interrupt signal, cleaning resources...")
			hnd.Close()

			glog.V(1).Infof("completed cleaning up resources")
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
