/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/mauleyzaola/maupod/src/server/maupod/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/psql"
	"github.com/spf13/viper"

	"github.com/mauleyzaola/maupod/src/server/maupod/pkg/api"
	"github.com/spf13/cobra"
)

const maupodDbName = "maupod"

// restapiCmd represents the restapi command
var restapiCmd = &cobra.Command{
	Use:   "restapi",
	Short: "Starts the restful application",
	Long:  `restapi will start a web server which is listening to requests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := api.ParseConfiguration()
		if err != nil {
			return err
		}

		pgConn := c.PgConn
		dbConn := pgConn + " dbname=" + maupodDbName

		var output io.Writer
		db, err := helpers.ConnectPostgres(pgConn, c.Retries, c.Delay)
		if err != nil {
			return err
		}

		// create database if not exist
		log.Println("creating database if not exists")
		if err = psql.CreateDbIfNotExists(db, maupodDbName); err != nil {
			return err
		}
		if err = db.Close(); err != nil {
			return err
		}

		// create the connection with the actual database
		log.Println("trying to connect to named database")
		if db, err = helpers.ConnectPostgres(dbConn, c.Retries, c.Delay); err != nil {
			return err
		}

		// TODO: create instance of the api server based on the real parameters
		apiServer, err := api.NewApiServer(db)
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
				log.Println("received an interrupt signal, cleaning resources...")
				if err = server.Shutdown(ctx); err != nil {
					log.Println(err)
				}
				if err = db.Close(); err != nil {
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
	},
}

func init() {
	rootCmd.AddCommand(restapiCmd)
}
