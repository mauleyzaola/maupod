package api

import (
	"database/sql"

	schema "github.com/gorilla/Schema"
)

type ApiServer struct {
	decoder *schema.Decoder
	db      *sql.DB
}

func NewApiServer(db *sql.DB) (*ApiServer, error) {
	s := &ApiServer{
		db:      db,
		decoder: schema.NewDecoder(),
	}
	return s, nil
}
