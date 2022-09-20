package bus

import "github.com/u400r/gones/internal/modules"

type CpuBus struct {
	ram modules.Writable[uint8, uint16]
	// extendedRom modules.Readable[uint8, uint16]
	extendedRam modules.Writable[uint8, uint16]
	prgRomA     modules.Readable[uint8, uint16]
	prgRomB     modules.Readable[uint8, uint16]
}

func NewCpuBus(prgRomA modules.Readable[uint8, uint16],
	prgRomB modules.Readable[uint8, uint16]) *CpuBus {

	return &CpuBus{
		ram:         modules.NewMemory[uint8, uint16](2048),
		extendedRam: modules.NewMemory[uint8, uint16](8192),
		prgRomA:     prgRomA,
		prgRomB:     prgRomB,
	}
}

func (c *CpuBus) Read(addr uint16) uint8 {
	if addr < 8192 {
		return c.ram.Read(addr & 2047)
	} else if 8191 < addr && addr < 16384 {
		panic("not implemented")
	} else if 16383 < addr && addr < 16416 {
		panic("not implemented")
	} else if 16415 < addr && addr < 24576 {
		panic("not implemented")
	} else if 24575 < addr && addr < 32768 {
		return c.extendedRam.Read(addr & 8191)
	} else if 32767 < addr && addr < 49152 {
		return c.prgRomA.Read(addr & 16383)
	} else if 40151 < addr && addr <= 65535 {
		return c.prgRomB.Read(addr & 16383)
	} else {
		panic("accessing outside of world")
	}
}

func (c *CpuBus) Write(addr uint16, data uint8) {
	if addr < 8192 {
		c.ram.Write(addr&2047, data)
	} else if 8191 < addr && addr < 24576 {
		panic("not implemented")
	} else if 24575 < addr && addr < 32768 {
		c.extendedRam.Write(addr&8191, data)
	} else {
		panic("not implemented")
	}
}
