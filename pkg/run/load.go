package run

import (
	"fmt"

	"github.com/aziule/bodar/pkg/behaviour"
	"github.com/aziule/bodar/pkg/behaviour/http"
	"github.com/aziule/bodar/pkg/config"
)

// Loader interface to load a user config.
type Loader interface {
	Load(cfg *config.UserConfig) error
}

// ConfigLoader is a struct used to load behaviours.
type ConfigLoader struct {
	runner *Runner
}

// NewLoader creates a new ConfigLoader.
func NewLoader(runner *Runner) *ConfigLoader {
	return &ConfigLoader{
		runner: runner,
	}
}

// Load behaviours given a user config.
func (l *ConfigLoader) Load(cfg *config.UserConfig) error {
	l.loadDefaultBehaviours()

	for _, bCfg := range cfg.Behaviours {
		_, ok := l.runner.available[bCfg.Name]
		if !ok {
			return fmt.Errorf(`behaviour "%s" not found`, bCfg.Name)
		}

		l.runner.Use(bCfg.Name, bCfg.Params)
	}

	return nil
}

// LoadCustomBehaviour makes a custom behaviour available to the runner.
func (l *ConfigLoader) LoadCustomBehaviour(name string, factoryFunc behaviour.FactoryFunc) *ConfigLoader {
	l.runner.available[name] = factoryFunc
	return l
}

func (l *ConfigLoader) loadDefaultBehaviours() *ConfigLoader {
	l.runner.available[http.SimpleResponseBehaviourName] = http.NewSimpleResponseBehaviour
	return l
}
