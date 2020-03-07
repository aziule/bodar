package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/config"
	"github.com/aziule/bodar/pkg/log"
)

const (
	defaultReadTimeout       = 10
	defaultReadHeaderTimeout = 10
	defaultWriteTimeout      = 10
	defaultIdleTimeout       = 10
	defaultMaxHeaderBytes    = 128
)

// Server is generic interface for an HTTP server.
type Server interface {
	Run(behaviour behaviour.Behaviour, port int, handlerFunc http.HandlerFunc) error
}

// DefaultServer is a default implementation of an http.Server.
type DefaultServer struct {
	srv *http.Server
}

// NewDefaultServer creates a new DefaultServer.
func NewDefaultServer(cfg config.BehaviourConfig) (*DefaultServer, error) {
	readTimeout, err := cfg.IntOrDefault("read_timeout", defaultReadTimeout)
	if err != nil {
		return nil, err
	}

	readHeaderTimeout, err := cfg.IntOrDefault("read_header_timeout", defaultReadHeaderTimeout)
	if err != nil {
		return nil, err
	}

	writeTimeout, err := cfg.IntOrDefault("write_timeout", defaultWriteTimeout)
	if err != nil {
		return nil, err
	}

	idleTimeout, err := cfg.IntOrDefault("idle_timeout", defaultIdleTimeout)
	if err != nil {
		return nil, err
	}

	maxHeaderBytes, err := cfg.IntOrDefault("max_header_bytes", defaultMaxHeaderBytes)
	if err != nil {
		return nil, err
	}

	return &DefaultServer{
		srv: &http.Server{
			ReadTimeout:       time.Duration(readTimeout) * time.Second,
			ReadHeaderTimeout: time.Duration(readHeaderTimeout) * time.Second,
			WriteTimeout:      time.Duration(writeTimeout) * time.Second,
			IdleTimeout:       time.Duration(idleTimeout) * time.Second,
			MaxHeaderBytes:    maxHeaderBytes,
		},
	}, nil
}

// Run the server and handle requests using the given handler.
func (s *DefaultServer) Run(behaviour behaviour.Behaviour, port int, handlerFunc http.HandlerFunc) error {
	log.Infof(`serving behaviour "%s" on port %d`, behaviour.Name(), port)
	s.setAddr(port)
	s.srv.Handler = ChainMiddlewares(handlerFunc, LogRequestMiddleware)
	return s.srv.ListenAndServe()
}

func (s *DefaultServer) setAddr(port int) {
	s.srv.Addr = fmt.Sprintf(":%d", port)
}
