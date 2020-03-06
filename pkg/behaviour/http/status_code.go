package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/behaviour"
)

// StatusCodeBehaviourName name.
const StatusCodeBehaviourName = "http-status-code"

// StatusCodeBehaviourName is an HTTP-based behaviour that returns specific status codes.
type StatusCodeBehaviour struct {
	server     Server
	port       int
	statusCode int
}

// Name returns the behaviour's name.
func (s *StatusCodeBehaviour) Name() string {
	return StatusCodeBehaviourName
}

// Run runs the HTTP server and serves the behaviour.
func (s *StatusCodeBehaviour) Run() error {
	return s.server.Run(s, s.port, s.handleRequest)
}

func (s *StatusCodeBehaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.statusCode)
	fmt.Println(s.statusCode)
}

// NewStatusCodeBehaviour creates a new StatusCodeBehaviour.
func NewStatusCodeBehaviour(cfg behaviour.Config) (behaviour.Behaviour, error) {
	server, ok := cfg["server"]
	if !ok {
		return nil, errors.New("missing config")
	}

	srv, ok := server.(Server)
	if !ok {
		return nil, fmt.Errorf(`invalid type found for config "server": Server expected, %T found`, server)
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
		server:     srv,
		port:       port,
		statusCode: statusCode,
	}
	return b, nil
}
