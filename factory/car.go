package factory

import "log"

type CarEngine struct{}

func (c CarEngine) Assemble() {
	log.Println("Assemble components for car engine")
}
