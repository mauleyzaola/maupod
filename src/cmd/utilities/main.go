package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/protos"

	_ "github.com/lib/pq"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
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

	// this is probably running from the host, not within a docker compose network
	const natsURL = "nats://localhost:4244"
	const pgConn = "user=postgres host=localhost port=5499 password=nevermind sslmode=disable dbname=maupod"

	nc, err := broker.ConnectNATS(natsURL, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	log.Println("successfully connected to NATS")

	delay := time.Second * time.Duration(config.Delay)
	if err = broker.RestAPIPing(nc, int(config.Retries), delay, delay); err != nil {
		return err
	}

	db, err := dbdata.ConnectPostgres(pgConn, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return err
	}
	defer func() {
		nc.Close()
		if err = db.Close(); err != nil {
			log.Println(err)
		}
	}()

	// TODO: make this better with cobra, for now I only need one command to be executed
	if err = cmdCheckFilesExist(db); err != nil {
		panic(err)
	}

	return nil
}

func cmdCheckFilesExist(conn boil.ContextExecutor) error {
	ctx := context.Background()
	rows, err := loadAllMedia(ctx, conn)
	if err != nil {
		return err
	}
	for _, row := range rows {
		if err = checkFileExist(row); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func loadAllMedia(ctx context.Context, conn boil.ContextExecutor) ([]*protos.Media, error) {
	store := dbdata.MediaStore{}
	var cols = orm.MediumColumns
	rows, err := store.List(ctx, conn, dbdata.MediaFilter{
		QueryFilter: dbdata.QueryFilter{
			Sort:      cols.Location,
			Direction: "asc",
		},
	}, nil)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func checkFileExist(m *protos.Media) error {
	var location = paths.MediaFullPathAudioFile(m.Location)
	_, err := os.Stat(location)
	return err
}
