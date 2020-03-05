package run

import (
	"context"
	"fmt"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/behaviour/http"
	"github.com/aziule/bodar/pkg/log"
)

// Runner is responsible for running the provided behaviours against a list of available behaviours.
type Runner struct {
	available map[string]behaviour.FactoryFunc
	enabled   map[string]map[string]interface{}
}

func (r *Runner) WithDefaultStrategies() *Runner {
	r.available = map[string]behaviour.FactoryFunc{
		http.EmptyBodyBehaviourName: http.NewEmptyBodyBehaviour,
	}
	return r
}

// Use defines what behaviour we want to use with specific config parameters.
func (r *Runner) Use(behaviour string, cfg map[string]interface{}) *Runner {
	if r.enabled == nil {
		r.enabled = make(map[string]map[string]interface{})
	}
	log.Infof(`adding behaviour "%s" to the list of desired behaviours`, behaviour)
	r.enabled[behaviour] = cfg
	return r
}

// Run tries to run the used behaviours using the registered behaviour factory funcs.
func (r *Runner) Run(ctx context.Context) error {
	log.Info("starting runner with the following behaviours:")

	ctx, cancel := context.WithCancel(ctx)

	for behaviour := range r.available {
		if _, ok := r.enabled[behaviour]; ok {
			log.Infof("\t+ %s", behaviour)
			continue
		}

		log.Infof("\t  %s", behaviour)
	}

	chErr := make(chan error)
	defer close(chErr)

	for name, cfg := range r.enabled {
		go func(name string, cfg map[string]interface{}) {
			chErr <- r.runBehaviour(name, cfg)
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

func (r *Runner) runBehaviour(name string, cfg map[string]interface{}) error {
	foundFactoryFunc, ok := r.available[name]
	if !ok {
		return fmt.Errorf(`behaviour "%s" not found`, name)
	}

	behaviour, err := foundFactoryFunc(cfg)
	if err != nil {
		return fmt.Errorf(`error creating behaviour "%s": %v`, name, err)
	}

	log.Infof(`running behaviour "%s"`, name)
	err = behaviour.Run()
	if err != nil {
		return fmt.Errorf(`error running behaviour "%s": %v`, name, err)
	}

	return nil
}
