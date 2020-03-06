package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/behaviour"
)

// EmptyBodyBehaviourName name.
const EmptyBodyBehaviourName = "http-empty-body"

// EmptyBodyBehaviour is an HTTP-based behaviour that always return empty bodies.
type EmptyBodyBehaviour struct {
	server Server
	port   int
}

// Name returns the behaviour's name.
func (s *EmptyBodyBehaviour) Name() string {
	return EmptyBodyBehaviourName
}

// Run runs the HTTP server and serves the behaviour.
func (s *EmptyBodyBehaviour) Run() error {
	return s.server.Run(s, s.port, s.handleRequest)
}

func (s *EmptyBodyBehaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
}

// NewEmptyBodyBehaviour creates a new EmptyBodyBehaviour.
func NewEmptyBodyBehaviour(cfg behaviour.Config) (behaviour.Behaviour, error) {
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

	b := &EmptyBodyBehaviour{
		server: srv,
		port:   port,
	}
	return b, nil
}
