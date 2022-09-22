package cpu

import (
	"fmt"

	"github.com/u400r/gones/internal/modules"
)

type Cpu struct {
	// registers
	aRegister              modules.WritableRegister[uint8]
	bRegister              modules.WritableRegister[uint8]
	xRegister              modules.WritableRegister[uint8]
	yRegister              modules.WritableRegister[uint8]
	programCounterRegister modules.Counter[uint16]
	statusRegister         modules.Flag[uint8]
	// stack
	stack *modules.Stack[uint8, uint8, uint16]
	// clock
	clock *Clock

	// outside component
	ram modules.Writable[uint8, uint16]

	// inner state
	instructionAddress uint16
	instructionOpecode uint8
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
		clock:                  &Clock{},
	}
	c.initDecoder()
	return c
}

type Clock struct {
	Cycles int
}

func (c *Clock) Tick() {
	c.Cycles += 1
}

func (c *Clock) GetCycles() int {
	return c.Cycles
}

type AddressingMode func(c *Cpu) *uint16

func (c *Cpu) Process() {
	opcode := c.fetch()
	op, mode := c.decode(opcode)
	if op == nil || mode == nil {
		panic("something went wrong")
	}

	addr := mode.GetAddress(c)
	op.Do(c, addr)
	c.programCounterRegister.Increment()
	c.printState(op, mode)

}

func (c *Cpu) printState(op Operatable, mode Addressing) {
	first, second := mode.GetOperandBytes(c)

	if first == nil {
		fmt.Printf("%X  %X        %v %v\n", c.instructionAddress, c.instructionOpecode, op.Opecode(), c.clock.Cycles)

	} else {
		if second == nil {
			fmt.Printf("%4X  %2X %2X     %v %v\n", c.instructionAddress, c.instructionOpecode, *first, op.Opecode(), c.clock.Cycles)

		} else {
			fmt.Printf("%4X  %2X %2X %2X  %v %v\n", c.instructionAddress, c.instructionOpecode, *first, *second, op.Opecode(), c.clock.Cycles)
		}
	}

}

func (c *Cpu) IsWaitOneClock() bool {
	return true
}

func (c *Cpu) fetch() uint8 {
	c.instructionAddress = c.programCounterRegister.Read()
	c.instructionOpecode = c.ram.Read(c.instructionAddress)
	return c.instructionOpecode
}
