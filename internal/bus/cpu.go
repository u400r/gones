package bus

import (
	"fmt"

	"github.com/u400r/gones/internal/controller"
	"github.com/u400r/gones/internal/modules"
	"github.com/u400r/gones/internal/ppu"
)

type CpuBus struct {
	ram modules.Writable[uint8, uint16]
	// extendedRom modules.Readable[uint8, uint16]
	extendedRam modules.Writable[uint8, uint16]
	prgRomA     modules.Readable[uint8, uint16]
	prgRomB     modules.Readable[uint8, uint16]
	debug       bool
	ppu         *ppu.Ppu
	cpuClock    *modules.Clock
	controllerA *controller.StandardController
	controllerB *controller.StandardController
}

func NewCpuBus(prgRomA modules.Readable[uint8, uint16],
	prgRomB modules.Readable[uint8, uint16], ppu *ppu.Ppu,
	cpuClock *modules.Clock,
	controllerA *controller.StandardController,
	controllerB *controller.StandardController) *CpuBus {

	return &CpuBus{
		ram:         modules.NewMemory[uint8, uint16](2048),
		extendedRam: modules.NewMemory[uint8, uint16](8192),
		prgRomA:     prgRomA,
		prgRomB:     prgRomB,
		ppu:         ppu,
		debug:       false,
		cpuClock:    cpuClock,
		controllerA: controllerA,
		controllerB: controllerB,
	}
}

func (c *CpuBus) Read(addr uint16) uint8 {
	if addr < 0x2000 {
		data := c.ram.Read(addr & 0x7FF)
		if c.debug {
			fmt.Printf("read  -> %04X %02X\n", addr&0x7FF, data)
		}
		return data
	} else if 0x1FFF < addr && addr < 0x4000 {
		return c.ppu.Read(addr)
	} else if addr == 0x4016 {
		return c.controllerA.Get()
	} else if addr == 0x4017 {
		return c.controllerB.Get()
	} else if 0x3FFF < addr && addr < 0x4020 {
		fmt.Printf("read from %04X not implemented\n", addr)
		return 0x0
	} else if 0x401F < addr && addr < 0x6000 {
		fmt.Printf("read from %04X not implemented\n", addr)
		return 0x0
	} else if 0x5FFF < addr && addr < 0x8000 {
		return c.extendedRam.Read(addr & 0x1FFF)
	} else if 0x7FFF < addr && addr < 0xC000 {
		data := c.prgRomA.Read(addr & 0x3FFF)
		if c.debug {
			fmt.Printf("prog  -> %04X %02X\n", addr&0x3FFF, data)
		}
		return data
	} else if 0xBFFF < addr && addr <= 0xFFFF {
		data := c.prgRomB.Read(addr & 0x3FFF)
		if c.debug {
			fmt.Printf("prog  -> %04X %02X\n", addr&0x3FFF, data)
		}
		return data
	} else {
		panic("accessing outside of world")
	}
}

func (c *CpuBus) Write(addr uint16, data uint8) {
	if addr < 0x2000 {
		if c.debug {
			fmt.Printf("write -> %04X %02X\n", addr&0x7FF, data)
		}
		c.ram.Write(addr&0x7FF, data)
	} else if 0x1FFF < addr && addr < 0x4000 {
		c.ppu.Write(addr, data)
	} else if addr == 0x4014 {
		// start address from cpu ram
		startAddr := uint16(data) << 8
		endAddr := startAddr | 0xFF
		c.cpuClock.Tock()
		// start address to oam
		for addr := startAddr; startAddr <= endAddr; addr += 0x1 {
			d := c.Read(addr)
			c.cpuClock.Tock()
			c.ppu.Write(0x2004, d)
			c.cpuClock.Tock()
		}
		if c.cpuClock.Cycles&1 == 1 {
			c.cpuClock.Tock()
		}
	} else if addr == 0x4016 {
		c.controllerA.ChangeStrobe(data&0x1 == 0x1)
		c.controllerB.ChangeStrobe(data&0x1 == 0x1)
	} else if 0x3FFF < addr && addr < 0x6000 && addr != 0x4014 {
		fmt.Printf("write to %04X not implemented\n", addr)
	} else if 0x5FFF < addr && addr < 0x8000 {
		c.extendedRam.Write(addr&0x1FFF, data)
	} else {
		panic("not implemented")
	}
}
