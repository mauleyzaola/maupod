package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

func SetupRoutes(a *ApiServer, output io.Writer) http.Handler {
	baseRouter := mux.NewRouter()
	baseRouter.Use(helpers.MiddlewareCORS)
	http.Handle("/", baseRouter)
	// Unauthenticated calls go through this router
	publicRouter := baseRouter.PathPrefix("/public").Subrouter()

	// TODO: use middleware authentication
	// authenticated routes
	//apiRouter := baseRouter.PathPrefix(baseUrl + "/api").Subrouter()

	publicRouter.HandleFunc("/system/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "pong")
	}).Methods(http.MethodOptions, http.MethodGet)

	publicRouter.HandleFunc("/audio/upload", a.GlueHandler(a.AudioFileUpload)).Methods(http.MethodOptions, http.MethodPost)

	if output != nil {
		return handlers.CombinedLoggingHandler(output, baseRouter)
	}
	return baseRouter
}
