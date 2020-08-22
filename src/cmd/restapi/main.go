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

	_ "github.com/lib/pq"
	"github.com/mauleyzaola/maupod/src/cmd/restapi/pkg/api"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
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

	nc, err := broker.ConnectNATS(config.NatsUrl, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	log.Println("successfully connected to NATS")

	var db *sql.DB
	if db, err = dbdata.DbBootstrap(config); err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if err = db.Close(); err != nil {
			log.Println(err)
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
		Addr:    ":8000",
		Handler: api.SetupRoutes(apiServer, output),
		//ReadTimeout:  TODO,
		//WriteTimeout: TODO,
	}

	// file server for artwork files
	if config.ArtworkStore != nil {
		fileServer := http.FileServer(http.Dir(config.ArtworkStore.Location))
		go func() {
			const port = ":9000"
			log.Println("starting file server at " + port + " from: " + config.ArtworkStore.Location)
			log.Fatal(http.ListenAndServe(port, fileServer))
		}()
	}

	var hnd types.Broker
	hnd = NewMsgHandler(nc)
	if err = hnd.Register(); err != nil {
		return err
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
			log.Println("received an interrupt signal, cleaning resources...")
			if err = server.Shutdown(ctx); err != nil {
				log.Println(err)
			}
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
	log.Println("app will exit now")

	return nil
}
