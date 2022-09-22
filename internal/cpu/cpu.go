package cpu

import (
	"fmt"

	"github.com/u400r/gones/internal/modules"
)

type Cpu struct {
	aRegister              modules.WritableRegister[uint8]
	bRegister              modules.WritableRegister[uint8]
	xRegister              modules.WritableRegister[uint8]
	yRegister              modules.WritableRegister[uint8]
	programCounterRegister modules.Counter[uint16]
	statusRegister         modules.Flag[uint8]
	stack                  *modules.Stack[uint8, uint8, uint16]
	ram                    modules.Writable[uint8, uint16]
	clock                  Clock
}

func NewCpu(memory modules.Writable[uint8, uint16]) *Cpu {
	c := &Cpu{
		aRegister:              modules.NewRegister(uint8(0)),
		bRegister:              modules.NewRegister(uint8(0)),
		xRegister:              modules.NewRegister(uint8(0)),
		yRegister:              modules.NewRegister(uint8(0)),
		programCounterRegister: modules.NewRegister(uint16(0)),
		statusRegister:         modules.NewRegister(uint8(0)),
		stack:                  modules.NewStack(memory, uint8(0)),
		ram:                    memory,
		clock:                  &clock{},
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
	c.programCounterRegister.Increment()

}

func (c *Cpu) IsWaitOneClock() bool {
	return true
}

func (c *Cpu) fetch() uint8 {
	return c.ram.Read(c.programCounterRegister.Read())
}
