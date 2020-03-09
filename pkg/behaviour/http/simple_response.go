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
	// SimpleResponseBehaviourName name.
	SimpleResponseBehaviourName = "http-simple-response"

	defaultStatusCode  = http.StatusOK
	defaultBody        = ""
	defaultContentType = ""
	defaultDelay       = 0
)

// SimpleResponseBehaviour is an HTTP-based behaviour that can return responses and a status code.
type SimpleResponseBehaviour struct {
	description string
	server      Server
	port        int
	statusCode  int
	contentType string
	body        []byte
	delay       time.Duration
}

// Name returns the behaviour's name.
func (s *SimpleResponseBehaviour) Name() string {
	return SimpleResponseBehaviourName
}

// Description returns the behaviour's description.
func (s *SimpleResponseBehaviour) Description() string {
	return s.description
}

// Run the HTTP server and handle requests.
func (s *SimpleResponseBehaviour) Run() error {
	log.Infof(`serving "%s" behaviour "%s" on port %d`, s.Name(), s.description, s.port)
	return s.server.Run(s.port, s.handleRequest)
}

func (s *SimpleResponseBehaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
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

// NewSimpleResponseBehaviour creates a new SimpleResponseBehaviour.
func NewSimpleResponseBehaviour(cfg config.BehaviourConfig) (behaviour.Behaviour, error) {
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

	b := &SimpleResponseBehaviour{
		description: description,
		server:      server,
		port:        port,
		statusCode:  statusCode,
		contentType: contentType,
		body:        []byte(body),
		delay:       time.Duration(delay) * time.Second,
	}
	return b, nil
}
