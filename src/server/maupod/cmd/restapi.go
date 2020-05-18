package cmd

import (
	"context"
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/mauleyzaola/maupod/src/server/maupod/pkg/api"
	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// scannerCmd represents the restapi command
var restapiCmd = &cobra.Command{
	Use:   "restapi",
	Short: "Starts the restful application",
	Long:  `restapi will start a web server which is listening to requests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := rule.ConfigurationParse()
		if err != nil {
			return err
		}

		if err = rule.ConfigurationValidate(config); err != nil {
			return err
		}

		var db *sql.DB
		if db, err = data.DbBootstrap(config); err != nil {
			return err
		}
		defer func() {
			if err = db.Close(); err != nil {
				log.Println(err)
			}
		}()

		var output io.Writer

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
