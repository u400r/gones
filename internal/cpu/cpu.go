package cpu

import (
	"github.com/u400r/gones/internal/modules"
)

type Cpu struct {
	a              modules.WritableRegister[uint8]
	b              modules.WritableRegister[uint8]
	x              modules.WritableRegister[uint8]
	y              modules.WritableRegister[uint8]
	programCounter modules.Counter[uint16]
	stack          *modules.Stack[uint8, uint8, uint16]
	ram            modules.Writable[uint8, uint16]
	clock          Clock
	status         modules.Flag[uint8]
}

func NewCpu(memory *modules.Memory[uint8, uint16]) *Cpu {
	return &Cpu{
		a:              modules.NewRegister(uint8(0)),
		b:              modules.NewRegister(uint8(0)),
		x:              modules.NewRegister(uint8(0)),
		y:              modules.NewRegister(uint8(0)),
		programCounter: modules.NewRegister(uint16(0)),
		status:         modules.NewRegister(uint8(0)),
		stack:          modules.NewStack(memory, uint8(0)),
		ram:            memory,
		clock:          &clock{},
	}
}

type Clock interface {
	Tick()
}
type clock struct {
}

func (c *clock) Tick() {

}

type Operation func(c *Cpu, addr *uint16)
type AddressingMode func(c *Cpu) *uint16

func (c *Cpu) Process() {
	opcode := c.fetch()
	operation, mode := c.decode(opcode)
	addr := mode(c)
	operation(c, addr)
	c.programCounter.Increment()
}

func (c *Cpu) IsWaitOneClock() bool {
	return true
}

func (c *Cpu) fetch() uint16 {
	return c.programCounter.Read()
}

func (c *Cpu) decode(opecode uint16) (Operation, AddressingMode) {
	decoded := OpecodeMapping[opecode]
	return decoded.Operation, decoded.AddressingMode
}
