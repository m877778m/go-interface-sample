package zeon

import "fmt"

func PrintZeon() {
	fmt.Println("Sieg Zeon!")
}

type ZeonMS interface {
	AvailableInSpace() bool
	WithBeamWepon() bool
	GetCode() string
}

func Check(ms ZeonMS) bool {
	if ms.AvailableInSpace() && ms.WithBeamWepon() {
		return true
	}
	return false
}

type Rfp struct {
	Comment string
	Pickup  []string
	Budget  int
}

func (rfp *Rfp) GetCodes() []string {
	return rfp.Pickup
}
