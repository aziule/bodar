package http

import (
	"fmt"
	"net/http"

	"github.com/aziule/bodar/internal/app/bodar"
)

const (
	EmptyBodyStrategyName        = "http-empty-body"
	EmptyBodyStrategyDescription = "Return empty HTTP response body"
)

// EmptyBodyStrategy is an HTTP-based strategy that always return empty bodies.
type EmptyBodyStrategy struct {
	server Server
	port   int
}

// Name returns the strategy's name.
func (s *EmptyBodyStrategy) Name() string {
	return EmptyBodyStrategyName
}

// Description returns the strategy's description.
func (s *EmptyBodyStrategy) Description() string {
	return EmptyBodyStrategyDescription
}

// Run runs the HTTP server and serves the strategy.
func (s *EmptyBodyStrategy) Run() error {
	return s.server.Run(s, s.port, s.handleRequest)
}

func (s *EmptyBodyStrategy) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
}

// NewEmptyBodyStrategy creates a new EmptyBodyStrategy.
func newEmptyBodyStrategy(cfg map[string]interface{}) (bodar.Strategy, error) {
	return &EmptyBodyStrategy{
		server: cfg["server"].(Server),
		port:   cfg["port"].(int),
	}, nil
}

func init() {
	bodar.RegisterStrategyFactoryFunc(EmptyBodyStrategyName, newEmptyBodyStrategy)
}
