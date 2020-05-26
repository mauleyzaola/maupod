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

	baseRouter.HandleFunc("/system/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "pong")
	}).Methods(http.MethodOptions, http.MethodGet)

	baseRouter.HandleFunc("/audio/scan", a.GlueHandler(a.AudioScanPost)).Methods(http.MethodOptions, http.MethodPost)

	if output != nil {
		return handlers.CombinedLoggingHandler(output, baseRouter)
	}
	return baseRouter
}
