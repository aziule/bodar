package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Config is the generic struct used to start the service.
type Config struct {
	Servers struct {
		Http map[string]struct {
			ReadTimeout       int `yaml:"read_timeout_ms"`
			ReadHeaderTimeout int `yaml:"read_header_timeout_ms"`
			WriteTimeout      int `yaml:"write_timeout_ms"`
			IdleTimeout       int `yaml:"idle_timeout_ms"`
			MaxHeaderBytes    int `yaml:"max_header_bytes"`
		} `yaml:"http,omitempty"`
	} `yaml:"servers""`
	Behaviours struct {
		Http []struct {
			Type   string                 `yaml:"type"`
			Server string                 `yaml:"server"`
			Port   int                    `yaml:"port"`
			Params map[string]interface{} `yaml:"params,omitempty"`
		} `yaml:"http,omitempty"`
	} `yaml:"behaviours"`
}

// Parse tries to parse bytes into Config.
func Parse(data []byte) (*Config, error) {
	var cfg Config
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid YAML format: %v", err)
	}
	return &cfg, nil
}
