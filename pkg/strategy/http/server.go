package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aziule/bodar/pkg/log"
	"github.com/aziule/bodar/pkg/strategy"
)

// Server is generic interface for an HTTP server.
type Server interface {
	Run(strategy strategy.Strategy, port int, handlerFunc http.HandlerFunc) error
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
			ReadTimeout:       cfg.ReadTimeout,
			ReadHeaderTimeout: cfg.ReadHeaderTimeout,
			WriteTimeout:      cfg.WriteTimeout,
			IdleTimeout:       cfg.IdleTimeout,
			MaxHeaderBytes:    cfg.MaxHeaderBytes,
		},
	}
}

// Run starts the server and serves the given func.
func (s *DefaultServer) Run(strategy strategy.Strategy, port int, handlerFunc http.HandlerFunc) error {
	log.Infof(`serving strategy "%s" on port %d`, strategy.Name(), port)
	s.setAddr(port)
	s.srv.Handler = handlerFunc
	return s.srv.ListenAndServe()
}

func (s *DefaultServer) setAddr(port int) {
	s.srv.Addr = fmt.Sprintf(":%d", port)
}
