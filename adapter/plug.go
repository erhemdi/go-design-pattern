package adapter

import "log"

type Plug interface {
	SupplyThreePin()
}

type SingaporePlug struct{}

func (p SingaporePlug) SupplyThreePin() {
	log.Println("Singapore plug: supply three pin")
}

type IndonesiaPlug struct{}

func (p IndonesiaPlug) SupplyTwoPin() {
	log.Println("Indonesia plug: supply two pin")
}

type IndonesiaPlugAdapter struct {
	IndonesiaPlug
}

func NewIndonesiaPlugAdapter(plug IndonesiaPlug) IndonesiaPlugAdapter {
	return IndonesiaPlugAdapter{plug}
}

func (p IndonesiaPlugAdapter) SupplyThreePin() {
	p.IndonesiaPlug.SupplyTwoPin()
}
