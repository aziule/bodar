package bodar

import "sync"

var strategyFactoryFuncs sync.Map

// StrategyFactoryFunc is a factory func used to create Strategy implementations given config parameters.
type StrategyFactoryFunc func(cfg map[string]interface{}) (Strategy, error)

// Strategy is a generic interface used to run different testing scenarios.
type Strategy interface {
	Run() error
	Name() string
	Description() string
}

// RegisterStrategyFactoryFunc registers a StrategyFactoryFunc.
func RegisterStrategyFactoryFunc(name string, factoryFunc StrategyFactoryFunc) {
	strategyFactoryFuncs.Store(name, factoryFunc)
}
