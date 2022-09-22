package modules

type Readable[T ByteSignal, U ByteSignal] interface {
	Read(addr U) T
}

type Writable[T ByteSignal, U ByteSignal] interface {
	Write(addr U, data T)
	Readable[T, U]
}

type Memory[T ByteSignal, U ByteSignal] struct {
	data []T
	size U
}

func NewMemory[T ByteSignal, U ByteSignal](size U) *Memory[T, U] {
	return &Memory[T, U]{
		data: make([]T, size),
		size: size,
	}
}

func NewMemoryWith[T ByteSignal, U ByteSignal](data []T) *Memory[T, U] {
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
