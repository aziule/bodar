package bodar

import (
	"sync"

	"github.com/aziule/bodar/internal/app/log"
)

// Registry is a container centralising the strategies that need to run.
type Registry struct {
	strategies []Strategy
}

// Register registers a new strategy to the registry.
func (r *Registry) Register(s Strategy) *Registry {
	r.strategies = append(r.strategies, s)
	return r
}

// Run tries to run the registered strategies.
func (r *Registry) Run() error {
	log.Info("running registry")
	wg := sync.WaitGroup{}
	for _, s := range r.strategies {
		wg.Add(1)
		go func(s Strategy) {
			log.Infof(`running strategy "%s"`, s.Name())
			err := s.Run()
			if err != nil {
				log.Errorf(`error running strategy "%s": %v`, s.Name(), err)
			}
			wg.Done()
		}(s)
	}
	wg.Wait()
	return nil
}

// Strategy is a generic interface used to run different testing scenarios.
type Strategy interface {
	Run() error
	Name() string
	Description() string
}
