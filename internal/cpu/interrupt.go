package cpu

import (
	"log"

	"github.com/u400r/gones/internal/modules"
)

func (c *Cpu) processRst() {
	c.statusRegister.Set(modules.INTERRUPT)
	low_address := c.ram.Read(0xFFFC)
	log.Print(low_address)
	high_address := c.ram.Read(0xFFFD)
	log.Print(high_address)
	address := uint16(high_address)<<8 + uint16(low_address)
	c.programCounterRegister.Write(address)
}

func (c *Cpu) processNmi() {
	c.statusRegister.Clear(modules.BLEAK)

	data := c.programCounterRegister.Read()
	high := uint8((data >> 8) & 255)
	low := uint8(data & 255)
	c.stack.Push(high)
	c.stack.Push(low)

	c.stack.Push(c.statusRegister.Read())

	c.statusRegister.Set(modules.INTERRUPT)
	low_address := c.ram.Read(0xFFFA)
	high_address := c.ram.Read(0xFFFB)

	address := uint16(high_address)<<8 + uint16(low_address)
	c.programCounterRegister.Write(address)

}

func (c *Cpu) processIrq() {

	if c.statusRegister.Get(modules.INTERRUPT) {
		return

	}

	c.statusRegister.Clear(modules.BLEAK)

	data := c.programCounterRegister.Read()
	high := uint8((data >> 8) & 255)
	low := uint8(data & 255)
	c.stack.Push(high)
	c.stack.Push(low)

	c.stack.Push(c.statusRegister.Read())

	c.statusRegister.Set(modules.INTERRUPT)
	low_address := c.ram.Read(0xFFFE)
	high_address := c.ram.Read(0xFFFF)

	address := uint16(high_address)<<8 + uint16(low_address)
	c.programCounterRegister.Write(address)

}

func (c *Cpu) PowerUp() {
	c.processRst()
}

func (c *Cpu) ProcessInterrupt() {
	if c.rstIn.Get() {
		c.processRst()
	} else if c.nmiIn.Get() {
		c.processNmi()
		c.nmiIn.Off()
	} else if c.irqIn.Get() {
		c.processIrq()
	}
}
