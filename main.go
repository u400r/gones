package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/u400r/gones/internal/bus"
	"github.com/u400r/gones/internal/controller"
	"github.com/u400r/gones/internal/cpu"
	"github.com/u400r/gones/internal/ines"
	"github.com/u400r/gones/internal/modules"
	"github.com/u400r/gones/internal/ppu"
)

func run() {
	nesFile := flag.String("nes-file", "", "")
	cpuDebug := flag.Bool("cpu-debug", false, "")
	ppuDebug := flag.Bool("ppu-debug", false, "")
	stepCpu := flag.Bool("step-cpu", false, "")
	stepPpu := flag.Bool("step-ppu", false, "")
	stepFrame := flag.Bool("step-frame", false, "")
	flag.Parse()
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 256, 240),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	cartrige := ines.ParseInes(*nesFile)
	memoryPpuBus := bus.NewPpuBus(cartrige.ChrRam, cartrige.Mirroring)
	rst := modules.NewBitSignal()
	nmi := modules.NewBitSignal()
	irq := modules.NewBitSignal()
	rst.Off()
	nmi.Off()
	irq.Off()
	cpuClock := modules.NewClock()
	ppuClock := modules.NewClock()
	controllerA := controller.NewController()
	controllerB := controller.NewController()
	ppu := ppu.NewPpu(memoryPpuBus, nmi, ppuClock, *ppuDebug, *stepPpu)
	memoryCpuBus := bus.NewCpuBus(cartrige.PrgRomA, cartrige.PrgRomB, ppu, cpuClock, controllerA, controllerB)
	cpu := cpu.NewCpu(memoryCpuBus, rst, nmi, irq, cpuClock, *cpuDebug, *stepCpu)
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
		if *stepFrame {
			bufio.NewScanner(os.Stdin).Scan()

		}
		controllerA.Toggle(controller.A, win.Pressed(pixelgl.KeyA))
		controllerA.Toggle(controller.B, win.Pressed(pixelgl.KeyB))
		controllerA.Toggle(controller.START, win.Pressed(pixelgl.KeyZ))
		controllerA.Toggle(controller.SELECT, win.Pressed(pixelgl.KeyX))
		controllerA.Toggle(controller.UP, win.Pressed(pixelgl.KeyUp))
		controllerA.Toggle(controller.DOWN, win.Pressed(pixelgl.KeyDown))
		controllerA.Toggle(controller.LEFT, win.Pressed(pixelgl.KeyLeft))
		controllerA.Toggle(controller.RIGHT, win.Pressed(pixelgl.KeyRight))
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
