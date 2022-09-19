package modules

type BitSignal interface {
	uint8 | uint16 | uint32 | uint64
}

type ReadableRegister[T BitSignal] interface {
	Read() T
}

type WritableRegister[T BitSignal] interface {
	Write(data T)
	ReadableRegister[T]
}

type Counter[T BitSignal] interface {
	Increment()
	Decrement()
	WritableRegister[T]
}

type Flag[T BitSignal] interface {
	WritableRegister[T]
	Get(nbit uint) bool
	Set(nbit uint)
	Clear(nbit uint)
	Change(nbit uint, flag bool)
}

type Register[T BitSignal] struct {
	data T
}

func NewRegister[T BitSignal](data T) *Register[T] {
	return &Register[T]{
		data: data,
	}
}

func (r *Register[T]) Read() T {
	return r.data
}

func (r *Register[T]) Write(data T) {
	r.data = data
}

func (r *Register[T]) Increment() {
	r.data += 1
}

func (r *Register[T]) Decrement() {
	r.data -= 1
}

func (r *Register[T]) Get(nbit uint) bool {
	return (r.data>>nbit)&0x1 == 0x1
}

func (r *Register[T]) Set(nbit uint) {
	r.data = r.data | 1<<nbit
}

func (r *Register[T]) Clear(nbit uint) {
	r.data = r.data & ^(1 << nbit)
}

func (r *Register[T]) Change(nbit uint, flag bool) {
	if flag {
		r.Set(nbit)
	} else {
		r.Clear(nbit)
	}
}

const (
	CARRY     uint = 0
	ZERO      uint = 1
	INTERRUPT uint = 2
	DECIMAL   uint = 3
	BLEAK     uint = 4
	OVERFLOW  uint = 6
	NEGATIVE  uint = 7
)
