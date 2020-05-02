package psql

import "database/sql"

//  ConnectPostgres should connect with or without a db schema
func ConnectPostgres(connStr string) (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}

func DatabaseExists(db *sql.DB, name string) (bool, error) {
	var query = `select count(*) from pg_database  where datname= $1`
	row := db.QueryRow(query, name)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

// CreateDb requires a connection without the dbname on it
func CreateDb(db *sql.DB, name string) error {
	var query = `create database ` + name
	_, err := db.Exec(query)
	return err
}

// CreateDbIfNotExists requires a connection without the dbname on it
func CreateDbIfNotExists(db *sql.DB, name string) error {
	ok, err := DatabaseExists(db, name)
	if err != nil {
		return err
	}
	if !ok {
		return CreateDb(db, name)
	}
	return nil
}
