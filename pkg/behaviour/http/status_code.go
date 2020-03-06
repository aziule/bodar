package http

import (
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/log"

	"github.com/aziule/bodar/pkg/config"

	"github.com/aziule/bodar/pkg/behaviour"
)

// StatusCodeBehaviourName name.
const StatusCodeBehaviourName = "http-status-code"

// StatusCodeBehaviour is an HTTP-based behaviour that returns specific status codes.
type StatusCodeBehaviour struct {
	server     Server
	port       int
	statusCode int
}

// Name returns the behaviour's name.
func (s *StatusCodeBehaviour) Name() string {
	return StatusCodeBehaviourName
}

// Run the HTTP server and handle requests.
func (s *StatusCodeBehaviour) Run() error {
	return s.server.Run(s, s.port, s.handleRequest)
}

func (s *StatusCodeBehaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Infof(`request received `)
	w.WriteHeader(s.statusCode)
}

// NewStatusCodeBehaviour creates a new StatusCodeBehaviour.
func NewStatusCodeBehaviour(cfg config.BehaviourConfig) (behaviour.Behaviour, error) {
	server, err := NewDefaultServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not create server: %v", err)
	}

	port, err := cfg.Int("port")
	if err != nil {
		return nil, err
	}

	statusCode, err := cfg.Int("status_code")
	if err != nil {
		return nil, err
	}

	b := &StatusCodeBehaviour{
		server:     server,
		port:       port,
		statusCode: statusCode,
	}
	return b, nil
}
