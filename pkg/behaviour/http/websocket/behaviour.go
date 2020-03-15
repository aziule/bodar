package websocket

import (
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/behaviour"
	apphttp "github.com/aziule/bodar/pkg/behaviour/http"
	"github.com/aziule/bodar/pkg/config"
	"github.com/aziule/bodar/pkg/log"
	"github.com/gorilla/websocket"
)

// BehaviourName name.
const BehaviourName = "websocket-default"

// Behaviour is a websocket-based behaviour.
type Behaviour struct {
	*behaviour.Base
	server apphttp.Server
	port   int
}

// Run runs the HTTP server and serves the websocket behaviour.
func (s *Behaviour) Run() error {
	log.Infof(`serving "%s" behaviour "%s" on port %d`, s.Name(), s.Description(), s.port)
	return s.server.Run(s.port, s.handleRequest)
}

func (s *Behaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
	var upgrader websocket.Upgrader

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("could not upgrade http connection: %v", err)
		return
	}
	defer func() {
		err := c.Close()
		if err != nil {
			log.Errorf("error closing the websocket connection: %v", err)
		}
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Errorf("read err: %v", err)
			break
		}
		log.Infof("received: %v", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Infof("written: %v", message)
			break
		}
	}
}

// NewBehaviour creates a new Behaviour.
func NewBehaviour(cfg config.BehaviourConfig) (behaviour.Behaviour, error) {
	base, err := behaviour.NewBase(BehaviourName, cfg)
	if err != nil {
		return nil, err
	}

	server, err := apphttp.NewDefaultServer(cfg)
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
