package bus

type Clock struct {
	Cycles int
	ch     chan bool
}

func NewClock() *Clock {
	return &Clock{
		Cycles: 0,
		ch:     make(chan bool),
	}
}

func (c *Clock) Tick() {
	c.Cycles += 1
	c.ch <- true
}

func (c *Clock) Tock() {
	<-c.ch
}

func (c *Clock) GetCycles() int {
	return c.Cycles
}
