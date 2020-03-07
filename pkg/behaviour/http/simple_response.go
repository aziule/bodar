package http

import (
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/config"
	"github.com/aziule/bodar/pkg/log"
)

// SimpleResponseBehaviourName name.
const SimpleResponseBehaviourName = "http-simple-response"

// SimpleResponseBehaviour is an HTTP-based behaviour that can return responses and a status code.
type SimpleResponseBehaviour struct {
	description string
	server      Server
	port        int
	statusCode  int
	contentType string
	body        []byte
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
	return s.server.Run(s, s.port, s.handleRequest)
}

func (s *SimpleResponseBehaviour) handleRequest(w http.ResponseWriter, r *http.Request) {
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

	description, err := cfg.StringOrDefault("description", "")
	if err != nil {
		return nil, err
	}

	port, err := cfg.Int("port")
	if err != nil {
		return nil, err
	}

	statusCode, err := cfg.IntOrDefault("status_code", http.StatusOK)
	if err != nil {
		return nil, err
	}

	body, err := cfg.StringOrDefault("body", "")
	if err != nil {
		return nil, err
	}

	contentType, err := cfg.StringOrDefault("content_type", "")
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
	}
	return b, nil
}
