package cpu

import (
	"fmt"

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

func NewCpu(memory modules.Writable[uint8, uint16]) *Cpu {
	c := &Cpu{
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
	c.initDecoder()
	return c
}

type Clock interface {
	Tick()
}
type clock struct {
}

func (c *clock) Tick() {

}

type AddressingMode func(c *Cpu) *uint16

func (c *Cpu) Process() {
	opcode := c.fetch()
	op, mode := c.decode(opcode)
	if op == nil || mode == nil {
		panic("something went wrong")
	}
	fmt.Printf("%v %v\n", op.Opecode(), mode.GetModeString())

	addr := mode.GetAddress(c)
	op.Do(c, addr)
	c.programCounter.Increment()

}

func (c *Cpu) IsWaitOneClock() bool {
	return true
}

func (c *Cpu) fetch() uint8 {
	return c.ram.Read(c.programCounter.Read())
}
