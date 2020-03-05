package strategy

// Strategy is a generic interface used to run different testing scenarios.
type Strategy interface {
	Run() error
	Name() string
}

// StrategyFactoryFunc is a factory func used to create strategies given config parameters.
type StrategyFactoryFunc func(cfg map[string]interface{}) (Strategy, error)
