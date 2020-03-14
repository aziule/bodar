package behaviour

import (
	"github.com/aziule/bodar/pkg/config"
)

// FactoryFunc is a factory func used to create behaviours given config parameters.
type FactoryFunc func(cfg config.BehaviourConfig) (Behaviour, error)

// Behaviour is a generic interface used to run different testing scenarios.
type Behaviour interface {
	Run() error
	ID() string
	Name() string
	Description() string
}

// Base struct implementation to mutualise common Behaviour properties and methods.
type Base struct {
	id          string
	name        string
	description string
}

// ID returns the behaviour's ID.
func (b *Base) ID() string {
	return b.id
}

// Name returns the behaviour's name.
func (b *Base) Name() string {
	return b.name
}

// Description returns the behaviour's description.
func (b *Base) Description() string {
	return b.description
}

// NewBase creates a new Base behaviour with mandatory parameters.
func NewBase(name string, cfg config.BehaviourConfig) (*Base, error) {
	id, err := cfg.String("id")
	if err != nil {
		return nil, err
	}

	description, err := cfg.String("description")
	if err != nil {
		return nil, err
	}

	return &Base{
		id:          id,
		name:        name,
		description: description,
	}, nil
}
