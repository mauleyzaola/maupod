package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	schema "github.com/gorilla/Schema"
	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

type ApiServer struct {
	config  *pb.Configuration
	decoder *schema.Decoder
	db      *sql.DB
	dm      *data.MediaStore
}

func NewApiServer(config *pb.Configuration, db *sql.DB) (*ApiServer, error) {
	s := &ApiServer{
		config:  config,
		db:      db,
		decoder: schema.NewDecoder(),
		dm:      &data.MediaStore{},
	}

	s.dm = &data.MediaStore{}

	return s, nil
}

func (a *ApiServer) GlueHandler(fn TransactionExecutor, promises ...PromiseExecutor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.JSONHandler(w, r, fn, promises...)
	}
}

func (a *ApiServer) JSONHandler(w http.ResponseWriter, r *http.Request, fn TransactionExecutor, promises ...PromiseExecutor) {
	var status int
	var result interface{}
	var err error
	if fn == nil {
		helpers.WriteJson(w, errors.New("missing parameter: fn"), http.StatusInternalServerError, nil)
		return
	}
	conn, err := a.db.Begin()
	if err != nil {
		helpers.WriteJson(w, err, http.StatusInternalServerError, nil)
		return
	}
	ctx := r.Context()
	defer func() {
		var localErr error
		if err != nil {
			localErr = conn.Rollback()
		} else {
			localErr = conn.Commit()
		}
		for _, p := range promises {
			if p != nil {
				p(PromiseExecutorParameter{
					ctx:    ctx,
					status: status,
					result: result,
					err:    err,
				})
			}
		}
		if localErr != nil {
			log.Println(localErr)
		}
	}()

	var params = TransactionExecutorParams{
		r:    r,
		w:    w,
		ctx:  ctx,
		conn: conn,
	}
	if status, result, err = fn(params); err != nil {
		helpers.WriteJson(w, err, status, nil)
		return
	}

	helpers.WriteJson(w, nil, status, result)
}
