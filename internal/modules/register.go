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
