package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// UserConfig is the generic struct used to start the service.
type UserConfig struct {
	Behaviours []struct {
		Name   string          `yaml:"name"`
		Params BehaviourConfig `yaml:",inline"`
	} `yaml:"behaviours"`
}

// Parse tries to parse bytes into UserConfig.
func Parse(data []byte) (*UserConfig, error) {
	var cfg UserConfig
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid YAML format: %v", err)
	}
	return &cfg, nil
}

// BehaviourConfig stores configuration values used to create behaviours.
type BehaviourConfig map[string]interface{}

// String returns the string representation of a value, if it exists.
func (c BehaviourConfig) String(name string) (string, error) {
	val, ok := c[name]
	if !ok {
		return "", c.errMissingConfig(name)
	}

	return c.stringVal(name, val)
}

// StringOrDefault returns the string value or a default value if it does not exist.
func (c BehaviourConfig) StringOrDefault(name, def string) (string, error) {
	val, ok := c[name]
	if !ok {
		return def, nil
	}

	return c.stringVal(name, val)
}

// Int returns the int representation of a value, if it exists.
func (c BehaviourConfig) Int(name string) (int, error) {
	val, ok := c[name]
	if !ok {
		return 0, c.errMissingConfig(name)
	}

	return c.intVal(name, val)
}

// IntOrDefault returns the int value or a default value if it does not exist.
func (c BehaviourConfig) IntOrDefault(name string, def int) (int, error) {
	val, ok := c[name]
	if !ok {
		return def, nil
	}

	return c.intVal(name, val)
}

func (c BehaviourConfig) intVal(name string, val interface{}) (int, error) {
	i, ok := val.(int)
	if !ok {
		return 0, c.errInvalidType(name, "int", val)
	}
	return i, nil
}

func (c BehaviourConfig) stringVal(name string, val interface{}) (string, error) {
	i, ok := val.(string)
	if !ok {
		return "", c.errInvalidType(name, "string", val)
	}
	return i, nil
}

func (c BehaviourConfig) errMissingConfig(name string) error {
	return fmt.Errorf(`missing config for "%s"`, name)
}

func (c BehaviourConfig) errInvalidType(name, expectedType string, value interface{}) error {
	return fmt.Errorf(`invalid type found for config "%s": %s expected, %T found`, name, expectedType, value)
}
