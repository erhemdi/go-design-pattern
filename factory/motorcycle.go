package factory

import "log"

type MotorcycleEngine struct{}

func (m MotorcycleEngine) Assemble() {
	log.Println("Assemble components for motorcycle engine")
}
