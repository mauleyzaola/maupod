package dbdata

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

const maupodDbName = "maupod"

// DbBootstrap creates the db if it doesn't exists and executes migrations
func DbBootstrap(config *pb.Configuration) (*sql.DB, error) {
	pgConn := config.PgConn
	dbConn := config.DbConn

	db, err := ConnectPostgres(pgConn, int(config.Retries), time.Second*time.Duration(config.Delay))
	if err != nil {
		return nil, err
	}

	// create database if not exist
	log.Println("creating database if not exists")
	if err = CreateDbIfNotExists(db, maupodDbName); err != nil {
		return nil, err
	}
	if err = db.Close(); err != nil {
		return nil, err
	}

	// create the connection with the actual database
	log.Println("trying to connect to named database")
	if db, err = ConnectPostgres(dbConn, int(config.Retries), time.Second*time.Duration(config.Delay)); err != nil {
		return nil, err
	}

	// run sql migrations
	var migrationDir = "/go/src/github.com/mauleyzaola/maupod/src/assets/db-migrations"
	count, err := MigrateDbFromPath(db, "postgres", migrationDir)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if count != 0 {
		log.Printf("executed: %v migrations", count)
	}
	return db, nil
}

func ConnectPostgres(dbConn string, retries int, delay time.Duration) (*sql.DB, error) {
	var db *sql.DB
	var err error

	if db, err = sql.Open("postgres", dbConn); err != nil {
		return nil, err
	}

	fn := func(retry int) bool {
		if err := db.Ping(); err != nil {
			return false
		}

		return true
	}
	log.Println("trying to establish connection with database using connection:", dbConn)
	ok, err := helpers.RetryFunc("connecting to postgres", retries, delay, fn)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("could not connect to postgres")
	}
	log.Println("successfully connected to postgres")
	return db, nil
}
