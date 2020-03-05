package http

import (
	"fmt"
	"net/http"

	"github.com/aziule/bodar/pkg/strategy"
)

// EmptyBodyStrategyName name.
const EmptyBodyStrategyName = "http-empty-body"

// EmptyBodyStrategy is an HTTP-based strategy that always return empty bodies.
type EmptyBodyStrategy struct {
	server Server
	port   int
}

// Name returns the strategy's name.
func (s *EmptyBodyStrategy) Name() string {
	return EmptyBodyStrategyName
}

// Run runs the HTTP server and serves the strategy.
func (s *EmptyBodyStrategy) Run() error {
	return s.server.Run(s, s.port, s.handleRequest)
}

func (s *EmptyBodyStrategy) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
}

// NewEmptyBodyStrategy creates a new EmptyBodyStrategy.
func NewEmptyBodyStrategy(cfg map[string]interface{}) (strategy.Strategy, error) {
	s := &EmptyBodyStrategy{
		server: cfg["server"].(Server),
		port:   cfg["port"].(int),
	}
	return s, nil
}
