package main

import (
	"github.com/u400r/gones/internal/bus"
	"github.com/u400r/gones/internal/cpu"
	"github.com/u400r/gones/internal/ines"
	"github.com/u400r/gones/internal/modules"
)

func main() {
	fileName := "/home/ukawa/nestest.nes"
	cartrige := ines.ParseInes(fileName)
	memoryBus := bus.NewCpuBus(cartrige.PrgRomA, cartrige.PrgRomB)
	rst := modules.BitSignal{}
	nmi := modules.BitSignal{}
	irq := modules.BitSignal{}
	rst.Off()
	nmi.Off()
	irq.Off()
	clock := bus.NewClock()
	cpu := cpu.NewCpu(memoryBus, rst, nmi, irq, clock)
	go cpu.Start()

	for {
		clock.Tick()
	}
}
