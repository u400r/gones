package ines

import (
	"os"

	"github.com/u400r/gones/internal/cartridge"
	"github.com/u400r/gones/internal/modules"
)

func ParseInes(fileName string) *cartridge.Cartridge {
	// parsing ines format file
	//
	cartridge := &cartridge.Cartridge{}
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	formatName := make([]byte, 4)
	fp.Read(formatName)
	romSize := make([]byte, 1)
	fp.Read(romSize)

	cartridge.PrgRomSize = uint16(romSize[0]) * 16384
	cartridge.ChrRomSize = uint16(romSize[0]) * 8192

	flag6 := make([]byte, 1)
	fp.Read(flag6)

	cartridge.Mirroring = flag6[0]>>7 == 1
	cartridge.HasExtendMemory = flag6[0]>>6&1 == 1
	cartridge.HasTrainer = flag6[0]>>5&1 == 1
	cartridge.IgnoreMirroring = flag6[0]>>4&1 == 1
	// lowerMapperNumber := flag6[0] & 0xf

	flag7 := make([]byte, 1)
	fp.Read(flag7)

	cartridge.VsUnisystem = flag7[0]>>7 == 1
	cartridge.Playchoice10 = flag7[0]>>6 == 1

	// upperMapperNumber := flag7[0] & 0xf

	unused := make([]byte, 8)
	fp.Read(unused)

	if cartridge.HasTrainer {
		trainer := make([]byte, 512)
		fp.Read(trainer)
		cartridge.TrainerRom = modules.NewMemoryWith[uint8, uint16](trainer)
	}

	prgRomA := make([]byte, 16384)
	fp.Read(prgRomA)
	cartridge.PrgRomA = modules.NewMemoryWith[uint8, uint16](prgRomA)

	if cartridge.PrgRomSize == 32768 {
		prgRomB := make([]byte, 16384)
		fp.Read(prgRomB)
		cartridge.PrgRomB = modules.NewMemoryWith[uint8, uint16](prgRomB)
	} else {
		cartridge.PrgRomB = cartridge.PrgRomA
	}

	chrRom := make([]byte, 8192)
	fp.Read(chrRom)
	cartridge.ChrRom = modules.NewMemoryWith[uint8, uint16](chrRom)

	return cartridge
}
