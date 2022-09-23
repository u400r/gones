package cpu

import (
	"fmt"

	"github.com/u400r/gones/internal/bus"
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
	clock *bus.Clock

	// outside component
	ram   modules.Writable[uint8, uint16]
	rstIn modules.BitSignal
	nmiIn modules.BitSignal
	irqIn modules.BitSignal

	// inner state
	instructionAddress uint16
	instructionOpecode uint8
	op                 Operatable
	mode               Addressing
	debug              bool
}

func NewCpu(memory modules.Writable[uint8, uint16],
	rst, nmi, irq modules.BitSignal, clock *bus.Clock) *Cpu {
	c := &Cpu{
		aRegister:              modules.NewRegister(uint8(0)),
		bRegister:              modules.NewRegister(uint8(0)),
		xRegister:              modules.NewRegister(uint8(0)),
		yRegister:              modules.NewRegister(uint8(0)),
		programCounterRegister: modules.NewRegister(uint16(0)),
		statusRegister:         modules.NewRegister(uint8(0x34)),
		stack:                  modules.NewStack(memory, uint8(0xFD)),
		ram:                    memory,
		clock:                  clock,
		rstIn:                  rst,
		nmiIn:                  nmi,
		irqIn:                  irq,
		debug:                  true,
	}
	c.initDecoder()
	return c
}

type AddressingMode func(c *Cpu) *uint16

func (c *Cpu) Process() {
	c.ProcessInterrupt()
	c.fetch()
	c.decode()

	addr := c.mode.GetAddress(c)
	c.op.Do(c, addr)
	c.programCounterRegister.Increment()
	c.clock.Tock()
	if c.debug {
		c.printState(addr)
	}

}

func (c *Cpu) Start() {
	c.PowerUp()

	for {
		c.Process()

	}
}

func (c *Cpu) printState(addr *uint16) {

	first, second := c.mode.GetOperandBytes(c)
	instruction := ""
	if first == nil {
		instruction = fmt.Sprintf("%04X  %02X        %v", c.instructionAddress, c.instructionOpecode, c.op.Opecode())

	} else {
		if second == nil {
			instruction = fmt.Sprintf("%04X  %02X %02X     %v", c.instructionAddress, c.instructionOpecode, *first, c.op.Opecode())

		} else {
			instruction = fmt.Sprintf("%04X  %02X %02X %02X  %v", c.instructionAddress, c.instructionOpecode, *first, *second, c.op.Opecode())
		}
	}

	registers := fmt.Sprintf("A:%02X X:%02X Y:%02X P:%02X SP:%02X", c.aRegister.Read(), c.xRegister.Read(), c.yRegister.Read(), c.statusRegister.Read(), c.stack.GetStackPointer())

	if addr != nil {
		fmt.Printf("%v %v %v %04X CYC:%v\n", instruction, registers, c.mode.GetModeString(), *addr, c.clock.Cycles)

	} else {
		fmt.Printf("%v %v %v %v CYC:%v\n", instruction, registers, c.mode.GetModeString(), addr, c.clock.Cycles)

	}

}

func (c *Cpu) fetch() {
	c.instructionAddress = c.programCounterRegister.Read()
	c.instructionOpecode = c.ram.Read(c.instructionAddress)

}
