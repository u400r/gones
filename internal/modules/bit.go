package modules

type BitSignal struct {
	state bool
}

func (b *BitSignal) On() {
	b.state = true
}

func (b *BitSignal) Off() {
	b.state = false
}

func (b *BitSignal) Toggle() {
	b.state = !b.state
}

func (b *BitSignal) Get() bool {
	return b.state
}
