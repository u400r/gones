package modules

import "sync"

type BitSignal struct {
	state bool
	m     sync.RWMutex
}

func NewBitSignal() *BitSignal {
	return &BitSignal{
		state: false,
		m:     sync.RWMutex{},
	}
}

func (b *BitSignal) On() {
	b.m.Lock()
	defer b.m.Unlock()
	b.state = true
}

func (b *BitSignal) Off() {
	b.m.Lock()
	defer b.m.Unlock()

	b.state = false
}

func (b *BitSignal) Toggle() {
	b.m.Lock()
	defer b.m.Unlock()

	b.state = !b.state
}

func (b *BitSignal) Get() bool {
	b.m.RLock()
	defer b.m.RUnlock()

	return b.state
}
