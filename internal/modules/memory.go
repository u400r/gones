package modules

type Readable[T BitSignal, U BitSignal] interface {
	Read(addr U) T
}

type Writable[T BitSignal, U BitSignal] interface {
	Write(addr U, data T)
	Readable[T, U]
}

type Memory[T BitSignal, U BitSignal] struct {
	data []T
	size U
}

func NewMemory[T BitSignal, U BitSignal](size U) *Memory[T, U] {
	return &Memory[T, U]{
		data: make([]T, size),
		size: size,
	}
}

func NewMemoryWith[T BitSignal, U BitSignal](data []T) *Memory[T, U] {
	size := U(len(data))
	return &Memory[T, U]{
		data: data,
		size: size,
	}

}

func (m *Memory[T, U]) Read(addr U) T {
	return m.data[addr]
}

func (m *Memory[T, U]) Write(addr U, data T) {
	m.data[addr] = data
}
