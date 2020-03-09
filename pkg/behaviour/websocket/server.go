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

type DefaultServer struct {
	srv      apphttp.Server
	upgrader websocket.Upgrader
}

func NewDefaultServer(cfg config.BehaviourConfig) *DefaultServer {
	srv, err := apphttp.NewDefaultServer(cfg)
	if err != nil {
		panic(err) // TODO: handle
	}

	var upgrader websocket.Upgrader

	return &DefaultServer{
		srv:      srv,
		upgrader: upgrader,
	}
}

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
	})
}
