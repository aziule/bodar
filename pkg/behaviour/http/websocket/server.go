package websocket

import (
	"net/http"

	apphttp "github.com/aziule/bodar/pkg/behaviour/http"
	"github.com/aziule/bodar/pkg/config"
	"github.com/aziule/bodar/pkg/log"
	"github.com/gorilla/websocket"
)

// Server is generic interface for an HTTP server.
type Server interface {
	Run(port int, handlerFunc http.HandlerFunc) error
}

// DefaultServer is a default implementation of a websocket server.
type DefaultServer struct {
	srv      apphttp.Server
	upgrader websocket.Upgrader
}

// NewDefaultServer creates a DefaultServer.
func NewDefaultServer(cfg config.BehaviourConfig) (*DefaultServer, error) {
	srv, err := apphttp.NewDefaultServer(cfg)
	if err != nil {
		return nil, err
	}

	var upgrader websocket.Upgrader

	return &DefaultServer{
		srv:      srv,
		upgrader: upgrader,
	}, nil
}

// Run the server and handle requests using the given handler.
func (s *DefaultServer) Run(port int, handlerFunc http.HandlerFunc) error {
	return s.srv.Run(port, func(w http.ResponseWriter, r *http.Request) {
		c, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Errorf("could not upgrade http connection: %v", err)
			return
		}
		defer c.Close()

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
		handlerFunc(w, r)
	})
}
