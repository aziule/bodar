package http

import (
	"context"
	"net/http"

	"github.com/aziule/bodar/pkg/log"
	"github.com/google/uuid"
)

const requestID = "request_id"

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
		log.Infof("http request %s received from %s %s %s", r.Context().Value(requestID), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}

// LogRequestMiddleware assigns an ID to a request.
func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), requestID, uuid.New().String())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
