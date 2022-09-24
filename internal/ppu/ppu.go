package ppu

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/u400r/gones/internal/bus"
	"github.com/u400r/gones/internal/modules"
)

type Ppu struct {
	ram               modules.Writable[uint8, uint16]
	nmiOut            *modules.BitSignal
	ppuCtrlRegister   modules.Flag[uint8]
	ppuMaskRegister   modules.Flag[uint8]
	ppuStatusRegister modules.Flag[uint8]
	oamAddrRegister   modules.WritableRegister[uint8]
	oamDataRegister   modules.WritableRegister[uint8]
	oamDMARegister    modules.WritableRegister[uint8]
	ppuScrolRegister  modules.WritableRegister[uint8]
	ppuAddrRegister   modules.WritableRegister[uint8]
	// internal registers
	vRegister           modules.Counter[uint16]
	tRegister           modules.Counter[uint16]
	fineXRegister       modules.WritableRegister[uint8]
	w                   modules.BitSignal
	patternLowRegister  modules.ShiftableRegister[uint8]
	patternHighRegister modules.ShiftableRegister[uint8]
	attributeRegister   modules.WritableRegister[uint8]
	clock               *bus.Clock
	x                   uint16
	y                   uint16
	image               *image.RGBA
	debug               bool
	step                bool
}

func NewPpu(memoryBus modules.Writable[uint8, uint16], nmiOut *modules.BitSignal, clock *bus.Clock, debug bool, step bool) *Ppu {
	return &Ppu{
		ram:               memoryBus,
		nmiOut:            nmiOut,
		ppuCtrlRegister:   modules.NewRegister[uint8](0x0),
		ppuMaskRegister:   modules.NewRegister[uint8](0x0),
		ppuStatusRegister: modules.NewRegister[uint8](0x0),
		ppuScrolRegister:  modules.NewRegister[uint8](0x0),
		ppuAddrRegister:   modules.NewRegister[uint8](0x0),
		oamAddrRegister:   modules.NewRegister[uint8](0x0),
		oamDataRegister:   modules.NewRegister[uint8](0x0),
		oamDMARegister:    modules.NewRegister[uint8](0x0),
		vRegister:         modules.NewRegister[uint16](0x0),
		tRegister:         modules.NewRegister[uint16](0x0),
		fineXRegister:     modules.NewRegister[uint8](0x0),

		w: modules.BitSignal{},

		patternLowRegister:  modules.NewRegister[uint8](0x0),
		patternHighRegister: modules.NewRegister[uint8](0x0),
		attributeRegister:   modules.NewRegister[uint8](0x0),
		clock:               clock,
		x:                   0x0,
		y:                   0x0,
		image:               image.NewRGBA(image.Rect(0, 0, 255, 239)),
		debug:               debug,
		step:                step,
	}

}

func (p *Ppu) Write(addr uint16, data uint8) {
	switch addr & 0x7 {
	case 0x0:
		p.ppuCtrlRegister.Write(data)
		nametableSelect := uint16(data&0b00000011) << 10
		t := p.tRegister.Read()
		t = (t & 0x73FF) | nametableSelect
		p.tRegister.Write(t)
	case 0x1:
		p.ppuMaskRegister.Write(data)
	case 0x2:
		p.ppuStatusRegister.Write(data)
	case 0x3:
		p.oamAddrRegister.Write(data)
	case 0x4:
		p.oamDataRegister.Write(data)
	case 0x5:
		p.ppuScrolRegister.Write(data)
		t := p.tRegister.Read()
		if p.w.Get() {
			ppuscroll_high2 := ((uint16(data) & 0b11000000) >> 6) << 8
			ppuscroll_mid3 := ((uint16(data) & 0b00111000) >> 3) << 5
			ppuscroll_low3 := (uint16(data) & 0b00000111) << 12
			t = (t & 0b000110000000000) | ppuscroll_high2 | ppuscroll_mid3 | ppuscroll_low3
			p.w.Off()
		} else {
			ppuscroll_high5 := (uint16(data) & 0b11111000) >> 3
			ppuscroll_low3 := data & 0b111
			p.fineXRegister.Write(ppuscroll_low3)
			t = (t & 0b111111111100000) | ppuscroll_high5
			p.w.On()
		}
		p.tRegister.Write(t)

	case 0x6:
		p.ppuAddrRegister.Write(data)
		t := p.tRegister.Read()
		if p.w.Get() {
			t = (t & 0xFF00) | uint16(data)
			p.vRegister.Write(t)
			p.w.Off()
		} else {
			t = (t & 0x00FF) | ((uint16(data) & 0x3f) << 8)
			p.w.On()
		}
		p.tRegister.Write(t)

	case 0x7:
		t := p.tRegister.Read()
		p.ram.Write(t&0x3fff, data)
		mode := p.ppuCtrlRegister.Get(modules.IncrementMode)
		if mode {
			p.tRegister.Increment32()
		} else {
			p.tRegister.Increment()
		}
	}
}

func (p *Ppu) Read(addr uint16) uint8 {
	var ret uint8
	switch addr & 0x7 {
	case 0x0:
		ret = p.ppuCtrlRegister.Read()
	case 0x1:
		ret = p.ppuMaskRegister.Read()
	case 0x2:
		ret = p.ppuStatusRegister.Read()
		p.w.Off()
	case 0x3:
		ret = p.oamAddrRegister.Read()
	case 0x4:
		ret = p.oamDataRegister.Read()
	case 0x5:
		ret = p.ppuScrolRegister.Read()
	case 0x6:
		ret = p.ppuAddrRegister.Read()
	case 0x7:
		t := p.tRegister.Read()
		ret = p.ram.Read(t & 0x3fff)
		mode := p.ppuCtrlRegister.Get(modules.IncrementMode)
		if mode {
			p.tRegister.Increment32()
		} else {
			p.tRegister.Increment()
		}
	}
	return ret
}

func (p *Ppu) getAttributeBit(attribute uint8, x, y uint16) uint8 {
	offset := ((x&0x30)>>4 + ((y&0x30)>>4)*2) * 2
	return (attribute >> offset) & 3
}

func (p *Ppu) fetchNametable() uint8 {
	nameAddress := 0x2000 | p.vRegister.Read()&0x0FFF

	p.tick()
	nametable := p.ram.Read(nameAddress)
	p.tick()
	return nametable
}

func (p *Ppu) fetchAttributetable() uint8 {
	v := p.vRegister.Read()
	attributeAddress := 0x23c0 | (v & 0x0C00) | ((v >> 4) & 0x38) | ((v >> 2) & 0x07)
	p.tick()
	attribute := p.ram.Read(attributeAddress)
	p.tick()
	return attribute
}

func (p *Ppu) fetchPattern(name uint8) (uint8, uint8) {
	bgtileSelect := p.ppuCtrlRegister.Get(modules.BgtileSelect)
	var baseAddress uint16
	if bgtileSelect {
		baseAddress = 0x1000
	} else {
		baseAddress = 0x0000
	}
	nameOffset := uint16(name) << 4
	patternAddress := baseAddress + nameOffset
	fine_y := (p.vRegister.Read() >> 12) & 0x7
	patternAddress = patternAddress + fine_y
	p.tick()
	patternLow := p.ram.Read(patternAddress)
	p.tick()

	patternAddress = patternAddress + 0x0008
	p.tick()
	patternHigh := p.ram.Read(patternAddress)

	return patternHigh, patternLow
}

func (p *Ppu) fetchPalette(attributeBit uint8) [4]uint8 {
	// i dont know why ppu can fetch palette
	// within one cycle because this involve
	// memory access. or, the address palette
	// placed is not memory, but register?
	var paletteBase uint16
	paletteBase = 0x3F00
	paletteOffset := (uint16(attributeBit) << 2) + 1
	paletteAddress := paletteBase + paletteOffset
	color0 := p.ram.Read(paletteBase)
	color1 := p.ram.Read(paletteAddress)
	color2 := p.ram.Read(paletteAddress + 1)
	color3 := p.ram.Read(paletteAddress + 2)
	//return 0x20, 0x00, 0x10, 0x20
	return [4]uint8{color0, color1, color2, color3}
}

func (p *Ppu) rendering() {
	if !(0 <= p.x && p.x < 256 && 0 <= p.y && p.y < 240) {
		return
	}

	x := p.x
	y := p.y
	patternLow := p.patternLowRegister.Read()
	patternHigh := p.patternHighRegister.Read()
	attribute := p.attributeRegister.Read()
	if p.x == 256 {
		// its true?
		return

	}
	attributeBit := p.getAttributeBit(attribute, x, y)

	palette := p.fetchPalette(attributeBit)

	fineX := p.fineXRegister.Read() & 0x7
	lowBit := (patternLow & (1 << (7 - fineX))) >> (7 - fineX)
	highBit := (patternHigh & (1 << (7 - fineX))) >> (7 - fineX)
	patternNumber := highBit*2 + lowBit

	nesColor := ColorMap[palette[patternNumber]]
	p.patternHighRegister.Left(false)
	p.patternLowRegister.Left(false)
	if p.debug {
		fmt.Println(x, y, attributeBit, palette, patternNumber, nesColor, patternLow, patternHigh, fineX)

	}
	p.image.Set(int(x), int(y), color.RGBA{nesColor[0], nesColor[1], nesColor[2], 255})
}

func (p *Ppu) incVramX() {
	v := p.vRegister.Read()
	if (v & 0x001F) == 31 {
		v = v & ^uint16(0x001F)
		v = v ^ 0x0400

	} else {
		v = v + 1
	}

	p.vRegister.Write(v)

}

func (p *Ppu) incVramY() {
	v := p.vRegister.Read()
	var y uint16
	if (v & 0x7000) != 0x7000 {
		v = v + 0x1000
	} else {
		v = v & ^uint16(0x7000)
		y = (v & 0x03E0) >> 5
		if y == 29 {
			y = 0
			v = v ^ 0x0800
		} else if y == 31 {
			y = 0
		} else {
			y = y + 1
		}
		v = (v & ^uint16(0x03E0)) | (y << 5)

	}
	p.vRegister.Write(v)

}

func (p *Ppu) syncXVt() {
	t := p.tRegister.Read()
	v := p.vRegister.Read()
	p.vRegister.Write(v & ^uint16(0x041F) | t&0x041F)
}

func (p *Ppu) syncYVt() {
	t := p.tRegister.Read()
	v := p.vRegister.Read()
	p.vRegister.Write(v & ^uint16(0x7BE0) | t&0x7BE0)
}

func (p *Ppu) tick() {
	p.clock.Tock()
	//fmt.Printf("X:%v Y:%v V:%v T:%v\n", p.x, p.y, p.vRegister.Read(), p.tRegister.Read())
	p.rendering()
	if p.x == 340 {
		p.x = 0
		if p.y == 261 {
			p.y = 0
			p.clock.Update()
		} else {
			p.y = p.y + 1
		}
	} else {
		p.x = p.x + 1
	}
}

func (p *Ppu) process() {
	// fetch name table
	if p.x == 1 && p.y == 241 {
		p.ppuStatusRegister.Set(modules.Vblank)
		nmi_enable := p.ppuCtrlRegister.Get(modules.NmiEnable)
		if nmi_enable {
			p.nmiOut.On()
		}
	}
	if p.x == 1 && p.y == 261 {
		p.ppuStatusRegister.Clear(modules.Vblank)
	}

	if ((0 < p.x && p.x < 257) || (320 < p.x && p.x < 337)) &&
		((0 <= p.y && p.y < 240) || p.y == 261) {
		name := p.fetchNametable()
		attribute := p.fetchAttributetable()
		high, low := p.fetchPattern(name)

		p.patternLowRegister.Load(low)
		p.patternHighRegister.Load(high)
		p.attributeRegister.Write(attribute)
		if p.x == 256 {
			p.incVramY()
		}
		p.incVramX()
		if p.step {
			bufio.NewScanner(os.Stdin).Scan()

		}
	}

	if p.x == 257 {
		p.syncXVt()
	}
	if 279 < p.x && p.x < 305 && p.y == 261 {
		p.syncYVt()
	}
	p.tick()
	if p.step {
		bufio.NewScanner(os.Stdin).Scan()

	}
}

func (p *Ppu) Start() {
	for {
		p.process()
	}
}

func (p *Ppu) GetImage() *image.RGBA {
	return p.image
}
