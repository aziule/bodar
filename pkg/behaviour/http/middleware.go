package http

import (
	"net/http"

	"github.com/aziule/bodar/pkg/log"
)

// Middleware func.
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// ChainMiddlewares helps chaining middlewares in a clean way.
func ChainMiddlewares(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	if len(middlewares) == 0 {
		return h
	}

	wrapped := h

	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	return wrapped
}

// LogRequestMiddleware logs incoming requests.
func LogRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("new request")
		next.ServeHTTP(w, r)
	}
}
