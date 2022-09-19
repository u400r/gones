package modules

type Stack[T BitSignal, U BitSignal] struct {
	stackPointer Counter[U]
	memory       Writable[T, U]
}

func NewStack[T BitSignal, U BitSignal](ram Writable[T, U], addr U) *Stack[T, U] {
	return &Stack[T, U]{
		stackPointer: NewRegister(addr),
		memory:       ram,
	}
}

func (s *Stack[T, U]) Pop() T {
	s.stackPointer.Decrement()
	addr := s.stackPointer.Read()
	data := s.memory.read(addr)
	return data
}

func (s *Stack[T, U]) Push(data T) {
	addr := s.stackPointer.Read()
	s.memory.write(addr, data)
	s.stackPointer.Increment()

}
