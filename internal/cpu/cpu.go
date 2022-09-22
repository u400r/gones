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
	ram   modules.Writable[uint8, uint16]
	rstIn modules.BitSignal
	nmiIn modules.BitSignal
	irqIn modules.BitSignal

	// inner state
	instructionAddress uint16
	instructionOpecode uint8
	debug              bool
}

func NewCpu(memory modules.Writable[uint8, uint16],
	rst, nmi, irq modules.BitSignal) *Cpu {
	c := &Cpu{
		aRegister:              modules.NewRegister(uint8(0)),
		bRegister:              modules.NewRegister(uint8(0)),
		xRegister:              modules.NewRegister(uint8(0)),
		yRegister:              modules.NewRegister(uint8(0)),
		programCounterRegister: modules.NewRegister(uint16(0)),
		statusRegister:         modules.NewRegister(uint8(0x34)),
		stack:                  modules.NewStack(memory, uint8(0xFD)),
		ram:                    memory,
		clock:                  &Clock{},
		rstIn:                  rst,
		nmiIn:                  nmi,
		irqIn:                  irq,
		debug:                  true,
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
	c.ProcessInterrupt()
	opcode := c.fetch()
	op, mode := c.decode(opcode)
	if op == nil || mode == nil {
		panic("something went wrong")
	}

	addr := mode.GetAddress(c)
	op.Do(c, addr)
	c.programCounterRegister.Increment()
	c.clock.Tick()
	if c.debug {
		c.printState(op, mode, addr)
	}

}

func (c *Cpu) printState(op Operatable, mode Addressing, addr *uint16) {
	first, second := mode.GetOperandBytes(c)
	instruction := ""
	if first == nil {
		instruction = fmt.Sprintf("%04X  %02X        %v", c.instructionAddress, c.instructionOpecode, op.Opecode())

	} else {
		if second == nil {
			instruction = fmt.Sprintf("%04X  %02X %02X     %v", c.instructionAddress, c.instructionOpecode, *first, op.Opecode())

		} else {
			instruction = fmt.Sprintf("%04X  %02X %02X %02X  %v", c.instructionAddress, c.instructionOpecode, *first, *second, op.Opecode())
		}
	}
	registers := fmt.Sprintf("A:%02X X:%02X Y:%02X P:%02X SP:%02X", c.aRegister.Read(), c.xRegister.Read(), c.yRegister.Read(), c.statusRegister.Read(), c.stack.GetStackPointer())

	if addr != nil {
		fmt.Printf("%v  %v %v %04X\n", instruction, registers, mode.GetModeString(), *addr)

	} else {
		fmt.Printf("%v  %v %v %v\n", instruction, registers, mode.GetModeString(), addr)

	}

}

func (c *Cpu) fetch() uint8 {
	c.instructionAddress = c.programCounterRegister.Read()
	c.instructionOpecode = c.ram.Read(c.instructionAddress)
	return c.instructionOpecode
}
