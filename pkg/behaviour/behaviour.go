package behaviour

import "github.com/aziule/bodar/pkg/config"

// Behaviour is a generic interface used to run different testing scenarios.
type Behaviour interface {
	Run() error
	Name() string
	Description() string
}

// FactoryFunc is a factory func used to create behaviours given config parameters.
type FactoryFunc func(cfg config.BehaviourConfig) (Behaviour, error)
