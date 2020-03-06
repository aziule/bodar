package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/log"
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
func NewDefaultServer(cfg behaviour.Config) (*DefaultServer, error) {
	readTimeout, err := cfg.Int("read_timeout")
	if err != nil {
		return nil, err
	}

	readHeaderTimeout, err := cfg.Int("read_header_timeout")
	if err != nil {
		return nil, err
	}

	writeTimeout, err := cfg.Int("write_timeout")
	if err != nil {
		return nil, err
	}

	idleTimeout, err := cfg.Int("idle_timeout")
	if err != nil {
		return nil, err
	}

	maxHeaderBytes, err := cfg.Int("max_header_bytes")
	if err != nil {
		return nil, err
	}

	return &DefaultServer{
		srv: &http.Server{
			ReadTimeout:       time.Duration(readTimeout) * time.Millisecond,
			ReadHeaderTimeout: time.Duration(readHeaderTimeout) * time.Millisecond,
			WriteTimeout:      time.Duration(writeTimeout) * time.Millisecond,
			IdleTimeout:       time.Duration(idleTimeout) * time.Millisecond,
			MaxHeaderBytes:    maxHeaderBytes,
		},
	}, nil
}

// Run starts the server and serves the given func.
func (s *DefaultServer) Run(behaviour behaviour.Behaviour, port int, handlerFunc http.HandlerFunc) error {
	log.Infof(`serving behaviour "%s" on port %d`, behaviour.Name(), port)
	s.setAddr(port)
	s.srv.Handler = handlerFunc
	return s.srv.ListenAndServe()
}

func (s *DefaultServer) setAddr(port int) {
	s.srv.Addr = fmt.Sprintf(":%d", port)
}
