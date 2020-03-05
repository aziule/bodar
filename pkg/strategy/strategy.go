package strategy

// Strategy is a generic interface used to run different testing scenarios.
type Strategy interface {
	Run() error
	Name() string
}

// FactoryFunc is a factory func used to create strategies given config parameters.
type FactoryFunc func(cfg map[string]interface{}) (Strategy, error)
