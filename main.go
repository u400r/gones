package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/u400r/gones/internal/bus"
	"github.com/u400r/gones/internal/cpu"
	"github.com/u400r/gones/internal/ines"
	"github.com/u400r/gones/internal/modules"
	"github.com/u400r/gones/internal/ppu"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 256, 240),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	fileName := "/home/ukawa/nestest.nes"
	cartrige := ines.ParseInes(fileName)
	memoryPpuBus := bus.NewPpuBus(cartrige.ChrRam, cartrige.Mirroring)
	rst := modules.BitSignal{}
	nmi := modules.BitSignal{}
	irq := modules.BitSignal{}
	rst.Off()
	nmi.Off()
	irq.Off()
	cpuClock := bus.NewClock()
	ppuClock := bus.NewClock()
	ppu := ppu.NewPpu(memoryPpuBus, nmi, ppuClock)
	memoryCpuBus := bus.NewCpuBus(cartrige.PrgRomA, cartrige.PrgRomB, ppu)
	cpu := cpu.NewCpu(memoryCpuBus, rst, nmi, irq, cpuClock)
	go cpu.Start()
	go ppu.Start()

	for !win.Closed() {
		go func() {
			for {
				ppuClock.Tick()
				ppuClock.Tick()
				ppuClock.Tick()
				cpuClock.Tick()
			}
		}()
		ppuClock.Sync()
		picture := pixel.PictureDataFromImage(ppu.GetImage())
		// FIXME It may be wrong that draw background as sprite
		sprite := pixel.NewSprite(picture, picture.Bounds())
		sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()
	}

}

func main() {

	pixelgl.Run(run)
}
