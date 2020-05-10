package cmd

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"

	_ "github.com/lib/pq"
	"github.com/mauleyzaola/maupod/src/server/maupod/pkg/api"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/psql"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const maupodDbName = "maupod"

// restapiCmd represents the restapi command
var restapiCmd = &cobra.Command{
	Use:   "restapi",
	Short: "Starts the restful application",
	Long:  `restapi will start a web server which is listening to requests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := api.ParseConfiguration()
		if err != nil {
			return err
		}

		if err = config.Validate(); err != nil {
			return err
		}

		pgConn := config.PgConn
		dbConn := pgConn + " dbname=" + maupodDbName

		var output io.Writer
		db, err := helpers.ConnectPostgres(pgConn, config.Retries, config.Delay)
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
		if db, err = helpers.ConnectPostgres(dbConn, config.Retries, config.Delay); err != nil {
			return err
		}
		defer func() {
			if err = db.Close(); err != nil {
				log.Println(err)
			}
		}()

		// run sql migrations
		count, err := data.MigrateDbFromPath(db, "postgres", filepath.Join("assets", "db-migrations"))
		if err != nil {
			log.Println(err)
			return err
		}
		if count != 0 {
			log.Printf("executed: %v migrations", count)
		}

		// TODO: create instance of the api server based on the real parameters
		apiServer, err := api.NewApiServer(config, db)
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
