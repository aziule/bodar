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
	// BehaviourName name.
	BehaviourName = "http-default"

	defaultStatusCode  = http.StatusOK
	defaultBody        = ""
	defaultContentType = ""
	defaultDelay       = 0
)

// Behaviour is an HTTP-based behaviour.
type Behaviour struct {
	*behaviour.Base
	server      Server
	port        int
	statusCode  int
	contentType string
	body        []byte
	delay       time.Duration
}

// Run the HTTP server and handle requests.
func (s *Behaviour) Run() error {
	log.Infof(`serving "%s" behaviour "%s" on port %d`, s.Name(), s.Description(), s.port)
	return s.server.Run(s.port, s.handleRequest)
}

func (s *Behaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
	time.Sleep(s.delay)

	if s.contentType != "" {
		w.Header().Set("Content-Type", s.contentType)
	}

	w.WriteHeader(s.statusCode)
	_, err := w.Write(s.body)
	if err != nil {
		log.Errorf(`error writing response for behaviour "%s" with status code "%d" and body "%s": %v`, s.Name(), s.statusCode, s.body, err)
	}
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

	statusCode, err := cfg.IntOrDefault("status_code", defaultStatusCode)
	if err != nil {
		return nil, err
	}

	body, err := cfg.StringOrDefault("body", defaultBody)
	if err != nil {
		return nil, err
	}

	contentType, err := cfg.StringOrDefault("content_type", defaultContentType)
	if err != nil {
		return nil, err
	}

	delay, err := cfg.IntOrDefault("delay", defaultDelay)
	if err != nil {
		return nil, err
	}

	b := &Behaviour{
		Base:        base,
		server:      server,
		port:        port,
		statusCode:  statusCode,
		contentType: contentType,
		body:        []byte(body),
		delay:       time.Duration(delay) * time.Second,
	}
	return b, nil
}
