package helpers

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

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
	ok, err := RetryFunc("connecting to postgres", retries, delay, fn)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("could not connect to postgres")
	}
	log.Println("successfully connected to postgres")
	return db, nil
}
