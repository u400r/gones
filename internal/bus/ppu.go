package bus

import "github.com/u400r/gones/internal/modules"

type PpuBus struct {
	ram       modules.Writable[uint8, uint16]
	chrRam    modules.Writable[uint8, uint16]
	mirroring bool
}

func NewPpuBus(chrRam modules.Writable[uint8, uint16], mirroring bool) *PpuBus {
	return &PpuBus{
		ram:       modules.NewMemory[uint8, uint16](2048),
		chrRam:    chrRam,
		mirroring: mirroring,
	}
}

func (p *PpuBus) Read(addr uint16) uint8 {
	var data uint8
	if addr < 0x2000 {
		data = p.chrRam.Read(addr)
	} else if 0x1FFF < addr {
		if p.mirroring {
			data = p.ram.Read(addr & 0x7FF)
		} else {
			data = p.ram.Read(addr&0x3FF + addr&0x0800>>1)
		}
	}
	return data
}

func (p *PpuBus) Write(addr uint16, data uint8) {
	if addr < 0x2000 {
		p.chrRam.Write(addr, data)
	} else if 0x1FFF < addr {
		if p.mirroring {
			p.ram.Write(addr&0x7FF, data)
		} else {
			p.ram.Write(addr&0x3FF+(addr&0x800)>>1, data)
		}
	}
}
