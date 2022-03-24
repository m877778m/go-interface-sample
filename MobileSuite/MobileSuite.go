package mobilesuite

// 武器構造体
type Wepon struct {
	Weight   int
	Type     string
	Range    int
	Magazine int
	Output   string
}

func (w *Wepon) IsLongRange() bool {
	if w.Range > 100 {
		return true
	}
	return false
}

// MSスペック構造体
type Spec struct {
	Height        int
	Weignt        int
	MainWepon     *Wepon
	SubWepon      *Wepon
	PilotCapacity int
	InSpace       bool
	InSea         bool
}

func (spec *Spec) AvailavleInSpace() bool {
	return spec.InSpace
}
func (spec *Spec) AvailavleInSea() bool {
	return spec.InSea
}
