package helpers

import (
	"net/http"
	"strings"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// ChainMiddleware will receive h as the final http handler and optionally, a number of middlewares as m
// this function will execute all the middlewares one inside the next in reverse order
// example:
//    ChainMiddleware(myHttpHandler, myFirstMiddleware, mySecondMiddleware)
func ChainMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	chain := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		chain = m[i](chain)
	}
	return chain
}

func MiddlewareCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var allowedHeaders = []string{
		//	"Accept",
		//	"Access-Control-Allow-Origin",
		//	"api_key",
		//	"Authorization",
		//	"Content-Type",
		//}
		var allowedMethods = []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
			http.MethodPatch,
			http.MethodOptions,
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ","))
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	}
}
