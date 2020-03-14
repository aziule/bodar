package websocket

import (
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/config"
	"github.com/aziule/bodar/pkg/log"
)

const (
	// Behaviour name.
	BehaviourName = "websocket-behaviour"
)

// Behaviour is a websocket-based behaviour.
type Behaviour struct {
	description string
	server      Server
	port        int
}

// Name returns the behaviour's name.
func (s *Behaviour) Name() string {
	return BehaviourName
}

// Description returns the behaviour's description.
func (s *Behaviour) Description() string {
	return s.description
}

func (s *Behaviour) Run() error {
	log.Infof(`serving "%s" behaviour "%s" on port %d`, s.Name(), s.description, s.port)
	return s.server.Run(s.port, s.handleRequest)
}

func (s *Behaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Infof("websocket request")
}

// NewBehaviour creates a new Behaviour.
func NewBehaviour(cfg config.BehaviourConfig) (behaviour.Behaviour, error) {
	server, err := NewDefaultServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not create server: %v", err)
	}

	description, err := cfg.String("description")
	if err != nil {
		return nil, err
	}

	port, err := cfg.Int("port")
	if err != nil {
		return nil, err
	}

	b := &Behaviour{
		description: description,
		server:      server,
		port:        port,
	}
	return b, nil
}
