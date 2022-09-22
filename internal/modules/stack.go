package modules

type Stack[T ByteSignal, U ByteSignal, V ByteSignal] struct {
	stackPointer Counter[U]
	memory       Writable[T, V]
}

func NewStack[T ByteSignal, U ByteSignal, V ByteSignal](ram Writable[T, V], addr U) *Stack[T, U, V] {
	return &Stack[T, U, V]{
		stackPointer: NewRegister(addr),
		memory:       ram,
	}
}

func (s *Stack[T, U, V]) Pop() T {
	s.stackPointer.Decrement()
	addr := s.stackPointer.Read()
	data := s.memory.Read(V(addr))
	return data
}

func (s *Stack[T, U, V]) Push(data T) {
	addr := s.stackPointer.Read()
	s.memory.Write(V(addr), data)
	s.stackPointer.Increment()

}

func (s *Stack[T, U, V]) GetStackPointer() U {
	return s.stackPointer.Read()
}

func (s *Stack[T, U, V]) SetStackPointer(addr U) {
	s.stackPointer.Write(addr)
}
