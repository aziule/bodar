package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aziule/bodar/internal/app/bodar"
	"github.com/aziule/bodar/internal/app/log"
)

// Server is generic interface for an HTTP server.
type Server interface {
	Run(strat bodar.Strategy, hf http.HandlerFunc) error
}

// DefaultServer is a default implementation of an http.Server.
type DefaultServer struct {
	srv *http.Server
}

// DefaultServerConfig is the config required to create a DefaultServer.
type DefaultServerConfig struct {
	Port              int
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

// NewDefaultServer creates a new DefaultServer.
func NewDefaultServer(cfg DefaultServerConfig) *DefaultServer {
	return &DefaultServer{
		srv: &http.Server{
			Addr:              fmt.Sprintf(":%d", cfg.Port),
			ReadTimeout:       cfg.ReadTimeout,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
			MaxHeaderBytes:    cfg.MaxHeaderBytes,
		},
	}
}

// Run starts the server and serves the given func.
func (s *DefaultServer) Run(strat bodar.Strategy, hf http.HandlerFunc) error {
	log.Infof(`serving strategy "%s" on "%s"`, strat.Name(), s.srv.Addr)
	s.srv.Handler = hf
	return s.srv.ListenAndServe()
}
