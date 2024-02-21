package cell

type Core struct {
	duration chan float64
	tick     chan struct{}
	antenna  chan struct{}
}

func NewCore() Core {
	return Core{
		duration: make(chan float64, 1),
		tick:     make(chan struct{}, 1),
		antenna:  make(chan struct{}),
	}
}
