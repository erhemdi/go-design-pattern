package factory

import (
	"errors"
)

const (
	Car        = 1
	Motorcycle = 2
)

type Engine interface {
	Assemble()
}

func GetEngine(engineType int) (Engine, error) {
	switch engineType {
	case Car:
		return CarEngine{}, nil
	case Motorcycle:
		return MotorcycleEngine{}, nil
	default:
		return nil, errors.New("engine type not found")
	}
}

func AssemblingEngine(engine Engine) {
	if engine == nil {
		return
	}

	engine.Assemble()
}
