package bodar

import (
	"sync"

	"github.com/aziule/bodar/internal/app/log"
)

// Registry is a container centralising the strategies that will run.
type Registry struct {
	strategies map[string]map[string]interface{}
}

// Use defines what strategy we want to use with what config parameters.
func (r *Registry) Use(strategy string, cfg map[string]interface{}) *Registry {
	if r.strategies == nil {
		r.strategies = make(map[string]map[string]interface{})
	}
	log.Infof(`adding strategy "%s" to the list of desired strategies`, strategy)
	r.strategies[strategy] = cfg
	return r
}

// Run tries to run the used strategies using the registered strategy factory funcs.
func (r *Registry) Run() error {
	log.Info("running desired strategies")
	wg := sync.WaitGroup{}
	for name, cfg := range r.strategies {
		wg.Add(1)
		go func(name string, cfg map[string]interface{}) {
			defer wg.Done()
			r.runStrategy(name, cfg)
		}(name, cfg)
	}
	wg.Wait()
	return nil
}

func (r *Registry) runStrategy(name string, cfg map[string]interface{}) {
	foundFactoryFunc, ok := strategyFactoryFuncs.Load(name)
	if !ok {
		log.Errorf(`strategy "%s" not found`, name)
		return
	}

	stratFactoryFunc, ok := foundFactoryFunc.(StrategyFactoryFunc)
	if !ok {
		log.Errorf(`strategy "%s" could not be converted to Strategy: %T found`, name, foundFactoryFunc)
		return
	}

	strategy, err := stratFactoryFunc(cfg)
	if err != nil {
		log.Errorf(`error creating strategy "%s": %v`, name, err)
		return
	}

	log.Infof(`running strategy "%s"`, name)
	err = strategy.Run()
	if err != nil {
		log.Errorf(`error running strategy "%s": %v`, name, err)
		return
	}
}
