package api

import (
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

	baseRouter.HandleFunc("/audio/scan", chainGlueCors(a.AudioScanPost)).Methods(http.MethodOptions, http.MethodPost)

	baseRouter.HandleFunc("/genres", chainGlueCors(a.GenresListGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/genres/artwork", chainGlueCors(a.GenreArtworksGet)).Methods(http.MethodOptions, http.MethodGet)

	baseRouter.HandleFunc("/ipc", chainGlueCors(a.IPCPost)).Methods(http.MethodOptions, http.MethodPost)

	baseRouter.HandleFunc("/events", chainFn(a.MediaEventsGet(), cors)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/events", chainFn(a.MediaEventsPost(), cors)).Methods(http.MethodOptions, http.MethodPost)

	baseRouter.HandleFunc("/file-browser/directory", chainGlueCors(a.DirectoryReadGet)).Methods(http.MethodOptions, http.MethodPost)

	baseRouter.HandleFunc("/media/{field}/distinct", chainGlueCors(a.DistinctListGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/media", chainGlueCors(a.MediaListGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/media/albums", chainGlueCors(a.AlbumViewListGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/media/{id}/spectrum", chainFn(a.MediaSpectrumGet(), cors)).Methods(http.MethodOptions, http.MethodGet)

	baseRouter.HandleFunc("/playlists/{id}", chainGlueCors(a.PlaylistGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/playlists", chainGlueCors(a.PlaylistsGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/playlists", chainGlueCors(a.PlaylistPost)).Methods(http.MethodOptions, http.MethodPost)
	baseRouter.HandleFunc("/playlists/{id}", chainGlueCors(a.PlaylistPut)).Methods(http.MethodOptions, http.MethodPut)
	baseRouter.HandleFunc("/playlists/{id}/items", chainGlueCors(a.PlaylistItems)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/playlists/{id}", chainGlueCors(a.PlaylistDelete)).Methods(http.MethodOptions, http.MethodDelete)
	baseRouter.HandleFunc("/playlists/{id}/items", chainGlueCors(a.PlaylistItemPost)).Methods(http.MethodOptions, http.MethodPost)
	baseRouter.HandleFunc("/playlists/{id}/items/{position}", chainGlueCors(a.PlaylistItemDelete)).Methods(http.MethodOptions, http.MethodDelete)
	baseRouter.HandleFunc("/playlists/{id}/items", chainGlueCors(a.PlaylistItemPut)).Methods(http.MethodOptions, http.MethodPut)

	baseRouter.HandleFunc("/providers/metadata/cover", chainGlueCors(a.ProviderMetaCoverGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/providers/metadata/cover/{album_identifier}", chainGlueCors(a.ProviderMetaCoverPut)).Methods(http.MethodOptions, http.MethodPut)

	baseRouter.HandleFunc("/queue", chainGlueCors(a.QueueGet)).Methods(http.MethodOptions, http.MethodGet)
	baseRouter.HandleFunc("/queue", chainGlueCors(a.QueuePost)).Methods(http.MethodOptions, http.MethodPost)
	baseRouter.HandleFunc("/queue/{index}", chainGlueCors(a.QueueDelete)).Methods(http.MethodOptions, http.MethodDelete)

	baseRouter.HandleFunc("/system/ping", chainFn(handlerPing, cors)).Methods(http.MethodOptions, http.MethodGet)

	if output != nil {
		return handlers.CombinedLoggingHandler(output, baseRouter)
	}
	return baseRouter
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJson(w, nil, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "pong",
	})
}
