package websocket

import (
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/config"
	"github.com/aziule/bodar/pkg/log"
)

// BehaviourName name.
const BehaviourName = "websocket-default"

// Behaviour is a websocket-based behaviour.
type Behaviour struct {
	*behaviour.Base
	server Server
	port   int
}

// Run runs the HTTP server and serves the websocket behaviour.
func (s *Behaviour) Run() error {
	log.Infof(`serving "%s" behaviour "%s" on port %d`, s.Name(), s.Description(), s.port)
	return s.server.Run(s.port, s.handleRequest)
}

func (s *Behaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Infof("websocket request")
}

// NewBehaviour creates a new Behaviour.
func NewBehaviour(cfg config.BehaviourConfig) (behaviour.Behaviour, error) {
	base, err := behaviour.NewBase(BehaviourName, cfg)
	if err != nil {
		return nil, err
	}

	server, err := NewDefaultServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("could not create server: %v", err)
	}

	port, err := cfg.Int("port")
	if err != nil {
		return nil, err
	}

	b := &Behaviour{
		Base:   base,
		server: server,
		port:   port,
	}
	return b, nil
}
