package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
)

func SetupRoutes(a *ApiServer, output io.Writer) http.Handler {
	baseRouter := mux.NewRouter()

	// setup middlewares as pointers to functions
	var chainFn = helpers.ChainMiddleware
	var cors = helpers.MiddlewareCORS
	var chainGlueCors = func(fn TransactionExecutor, promises ...PromiseExecutor) http.HandlerFunc {
		return chainFn(a.GlueHandler(fn, promises...), cors)
	}

	http.Handle("/", baseRouter)

	baseRouter.HandleFunc("/system/ping", handlerPing).Methods(http.MethodOptions, http.MethodGet)

	baseRouter.HandleFunc("/audio/scan", chainGlueCors(a.AudioScanPost)).Methods(http.MethodOptions, http.MethodPost)
	baseRouter.HandleFunc("/artwork/scan", chainGlueCors(a.ArtworkScanPost)).Methods(http.MethodOptions, http.MethodPost)

	baseRouter.HandleFunc("/genres", chainGlueCors(a.GenresListGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/genres/artwork", chainGlueCors(a.GenreArtworksGet)).Methods(http.MethodOptions, http.MethodGet)

	baseRouter.HandleFunc("/ipc", chainGlueCors(a.IPCPost)).Methods(http.MethodOptions, http.MethodPost)

	baseRouter.HandleFunc("/media/{field}/distinct", chainGlueCors(a.DistinctListGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/media", chainGlueCors(a.MediaListGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/media/albums", chainGlueCors(a.AlbumViewListGet)).Methods(http.MethodOptions, http.MethodGet)

	baseRouter.HandleFunc("/queue/list", chainGlueCors(a.QueueListGet)).Methods(http.MethodOptions, http.MethodGet)

	if output != nil {
		return handlers.CombinedLoggingHandler(output, baseRouter)
	}
	return baseRouter
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "pong")
}
