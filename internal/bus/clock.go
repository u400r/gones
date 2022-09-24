package bus

type Clock struct {
	Cycles int
	clock  chan bool
	sync   chan bool
}

func NewClock() *Clock {
	return &Clock{
		Cycles: 0,
		clock:  make(chan bool),
		sync:   make(chan bool),
	}
}

func (c *Clock) Tick() {
	c.clock <- true
}

func (c *Clock) Tock() {
	c.Cycles += 1
	<-c.clock
}

func (c *Clock) Update() {
	c.sync <- true
}

func (c *Clock) Sync() {
	<-c.sync
}

func (c *Clock) GetCycles() int {
	return c.Cycles
}
