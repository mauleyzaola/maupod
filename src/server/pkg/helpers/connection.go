package helpers

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/data/psql"
	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
)

const maupodDbName = "maupod"

func DbBootstrap(config *domain.Configuration) (*sql.DB, error) {
	pgConn := config.PgConn
	dbConn := pgConn + " dbname=" + maupodDbName

	db, err := ConnectPostgres(pgConn, config.Retries, config.Delay)
	if err != nil {
		return nil, err
	}

	// create database if not exist
	log.Println("creating database if not exists")
	if err = psql.CreateDbIfNotExists(db, maupodDbName); err != nil {
		return nil, err
	}
	if err = db.Close(); err != nil {
		return nil, err
	}

	// create the connection with the actual database
	log.Println("trying to connect to named database")
	if db, err = ConnectPostgres(dbConn, config.Retries, config.Delay); err != nil {
		return nil, err
	}

	// run sql migrations
	count, err := MigrateDbFromPath(db, "postgres", filepath.Join("assets", "db-migrations"))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if count != 0 {
		log.Printf("executed: %v migrations", count)
	}
	return db, nil
}
