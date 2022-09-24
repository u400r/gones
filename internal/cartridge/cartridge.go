package cartridge

import "github.com/u400r/gones/internal/modules"

type Cartridge struct {
	PrgRomSize      uint16
	ChrRamSize      uint16
	InesVersion     string
	Mirroring       bool
	HasExtendMemory bool
	HasTrainer      bool
	IgnoreMirroring bool
	MapperNumber    uint16
	VsUnisystem     bool
	Playchoice10    bool

	PrgRomA modules.Readable[uint8, uint16]
	PrgRomB modules.Readable[uint8, uint16]
	ChrRam  modules.Writable[uint8, uint16]
	//extend_ram
	TrainerRom modules.Readable[uint8, uint16]
}
