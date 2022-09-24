package modules

import "sync"

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
	m    sync.RWMutex
}

func NewMemory[T ByteSignal, U ByteSignal](size U) *Memory[T, U] {
	return &Memory[T, U]{
		data: make([]T, size),
		size: size,
		m:    sync.RWMutex{},
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
	m.m.RLock()
	defer m.m.RUnlock()
	return m.data[addr]
}

func (m *Memory[T, U]) Write(addr U, data T) {
	m.m.Lock()
	defer m.m.Unlock()
	m.data[addr] = data
}
