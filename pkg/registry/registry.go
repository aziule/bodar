package registry

import (
	"context"
	"fmt"

	"github.com/aziule/bodar/pkg/strategy/http"

	"github.com/aziule/bodar/pkg/strategy"

	"github.com/aziule/bodar/pkg/log"
)

// Registry is a container centralising the strategies that will run.
type Registry struct {
	available map[string]strategy.StrategyFactoryFunc
	enabled   map[string]map[string]interface{}
}

func (r *Registry) WithDefaultStrategies() *Registry {
	r.available = map[string]strategy.StrategyFactoryFunc{
		http.EmptyBodyStrategyName: http.NewEmptyBodyStrategy,
		"2":                        http.NewEmptyBodyStrategy,
		"3":                        http.NewEmptyBodyStrategy,
		"4":                        http.NewEmptyBodyStrategy,
		"5":                        http.NewEmptyBodyStrategy,
	}
	return r
}

// Use defines what strategy we want to use with what config parameters.
func (r *Registry) Use(strategy string, cfg map[string]interface{}) *Registry {
	if r.enabled == nil {
		r.enabled = make(map[string]map[string]interface{})
	}
	log.Infof(`adding strategy "%s" to the list of desired strategies`, strategy)
	r.enabled[strategy] = cfg
	return r
}

// Run tries to run the used strategies using the registered strategy factory funcs.
func (r *Registry) Run(ctx context.Context) error {
	log.Info("starting registry with the following strategies:")

	ctx, cancel := context.WithCancel(ctx)

	for strategy := range r.available {
		if _, ok := r.enabled[strategy]; ok {
			log.Infof("\t+ %s", strategy)
			continue
		}

		log.Infof("\t  %s", strategy)
	}

	chErr := make(chan error)
	defer close(chErr)

	for name, cfg := range r.enabled {
		go func(name string, cfg map[string]interface{}) {
			chErr <- r.runStrategy(name, cfg)
		}(name, cfg)
	}

	select {
	case err := <-chErr:
		log.Error(err)
	case <-ctx.Done():
		break
	}

	cancel()

	return nil
}

func (r *Registry) runStrategy(name string, cfg map[string]interface{}) error {
	foundFactoryFunc, ok := r.available[name]
	if !ok {
		return fmt.Errorf(`strategy "%s" not found`, name)
	}

	strategy, err := foundFactoryFunc(cfg)
	if err != nil {
		return fmt.Errorf(`error creating strategy "%s": %v`, name, err)
	}

	log.Infof(`running strategy "%s"`, name)
	err = strategy.Run()
	if err != nil {
		return fmt.Errorf(`error running strategy "%s": %v`, name, err)
	}

	return nil
}
