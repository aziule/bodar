package behaviour

import (
	"fmt"
)

// Behaviour is a generic interface used to run different testing scenarios.
type Behaviour interface {
	Run() error
	Name() string
}

// FactoryFunc is a factory func used to create behaviours given config parameters.
type FactoryFunc func(cfg Config) (Behaviour, error)

// Config stores configuration values used to create behaviours.
type Config map[string]interface{}

// Int returns the int representation of a value, if it exists.
func (c Config) Int(name string) (int, error) {
	val, ok := c[name]
	if !ok {
		return 0, c.errMissingConfig(name)
	}

	valInt, ok := val.(int)
	if !ok {
		return 0, c.errInvalidType(name, "int", val)
	}

	return valInt, nil
}

func (c Config) errMissingConfig(name string) error {
	return fmt.Errorf(`missing config for "%s"`, name)
}

func (c Config) errInvalidType(name, expectedType string, value interface{}) error {
	return fmt.Errorf(`invalid type found for config "%s": %s expected, %T found`, name, expectedType, value)
}
