package main

import (
	"context"
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/simplelog"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"

	_ "github.com/lib/pq"
	"github.com/mauleyzaola/maupod/src/server/maupod/pkg/api"
	"github.com/mauleyzaola/maupod/src/server/pkg/data"
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
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
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

	nc, err := helpers.ConnectNATS(int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	logger.Info("successfully connected to NATS")

	var db *sql.DB
	if db, err = data.DbBootstrap(config); err != nil {
		return err
	}
	defer func() {
		if err = db.Close(); err != nil {
			logger.Error(err)
		}
	}()

	var output io.Writer

	// TODO: create instance of the api server based on the real parameters
	apiServer, err := api.NewApiServer(config, db, nc)
	if err != nil {
		return err
	}

	// TODO: allow more options, for now stdout is ok
	output = os.Stdout

	signalChan := make(chan os.Signal, 1)
	server := http.Server{
		Addr:    viper.GetString("port"),
		Handler: api.SetupRoutes(apiServer, output),
		//ReadTimeout:  TODO,
		//WriteTimeout: TODO,
	}

	go func() {
		log.Println("api serving from ", server.Addr)
		log.Fatal(server.ListenAndServe())
	}()

	ctx := context.Background()
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		cleanup := func() {
			logger.Info("received an interrupt signal, cleaning resources...")
			if err = server.Shutdown(ctx); err != nil {
				log.Println(err)
			}
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
	log.Println("app will exit now")

	return nil
}
