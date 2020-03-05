package behaviour

// Behaviour is a generic interface used to run different testing scenarios.
type Behaviour interface {
	Run() error
	Name() string
}

// FactoryFunc is a factory func used to create behaviours given config parameters.
type FactoryFunc func(cfg map[string]interface{}) (Behaviour, error)
