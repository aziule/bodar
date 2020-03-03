package http

import (
	"fmt"
	"net/http"
)

// EmptyBodyStrategy is an HTTP-based strategy that always return empty bodies.
type EmptyBodyStrategy struct {
	server Server
}

// EmptyBodyStrategyConfig contains the EmptyBodyStrategy config.
type EmptyBodyStrategyConfig struct {
	Server Server
}

// NewEmptyBodyStrategy creates a new EmptyBodyStrategy.
func NewEmptyBodyStrategy(cfg EmptyBodyStrategyConfig) *EmptyBodyStrategy {
	return &EmptyBodyStrategy{
		server: cfg.Server,
	}
}

// Name returns the strategy's name.
func (s *EmptyBodyStrategy) Name() string {
	return "http-empty-body"
}

// Description returns the strategy's description.
func (s *EmptyBodyStrategy) Description() string {
	return fmt.Sprintf("Return empty HTTP response body")
}

// Run runs the HTTP server and serves the strategy.
func (s *EmptyBodyStrategy) Run() error {
	return s.server.Run(s, s.handleRequest)
}

func (s *EmptyBodyStrategy) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
}
