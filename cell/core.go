package cell

type Core struct {
	duration chan float64
	tick     chan struct{}
	antenna  chan struct{}
}

func NewCore() *Core {
	return &Core{
		duration: make(chan float64, 1),
		tick:     make(chan struct{}, 1),
		antenna:  make(chan struct{}),
	}
}

func (c *Core) GetDuration() chan float64 {
	return c.duration
}
func (c *Core) GetAntenna() chan struct{} {
	return c.antenna
}
func (c *Core) GetTick() chan struct{} {
	return c.tick
}
