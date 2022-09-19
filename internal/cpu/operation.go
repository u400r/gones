package cpu

import "github.com/u400r/gones/internal/modules"

func OperationAdc(c *Cpu, addr *uint16) {
	c.clock.Tick()
	a := c.a.Read()
	memory := c.ram.Read(*addr)
	carryIn := c.status.Get(modules.CARRY)

	result, carryOut := modules.UnsignedAdd(a, memory, carryIn)
	_, overflow := modules.SignedAdd(int8(a), int8(memory), carryIn)
	c.a.Write(result)
	c.status.Change(modules.OVERFLOW, overflow)
	c.status.Change(modules.NEGATIVE, result>>7 == 1)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, carryOut)
}

func OperationSbc(c *Cpu, addr *uint16) {
	c.clock.Tick()
	a := c.a.Read()
	memory := c.ram.Read(*addr)
	carryIn := c.status.Get(modules.CARRY)

	result, carryOut := modules.UnsignedSub(a, memory, carryIn)
	_, overflow := modules.SignedSub(int8(a), int8(memory), carryIn)

	c.a.Write(result)
	c.status.Change(modules.OVERFLOW, overflow)
	c.status.Change(modules.NEGATIVE, result>>7 == 1)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, carryOut)
}

func OperationAnd(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := memory & a
	c.a.Write(result)

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationOra(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := memory | a
	c.a.Write(result)

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationEor(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := memory ^ a
	c.a.Write(result)

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationAsl(c *Cpu, addr *uint16) {
	c.clock.Tick()
	var data uint8
	if addr != nil {
		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {
		data = c.a.Read()
	}

	result := (data << 1) & 255
	carryOut := (data >> 7) == 1

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, carryOut)
	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.a.Write(result)
	}
}

func OperationLsr(c *Cpu, addr *uint16) {
	c.clock.Tick()
	var data uint8
	if addr != nil {
		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {
		data = c.a.Read()
	}
	result := data >> 1
	carryOut := data == 1

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, carryOut)
	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.a.Write(result)
	}
}

func OperationRol(c *Cpu, addr *uint16) {
	c.clock.Tick()

	var data uint8
	if addr != nil {
		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {

		data = c.a.Read()
	}
	var carryIn uint8
	if c.status.Get(modules.CARRY) {
		carryIn = 1
	} else {
		carryIn = 0
	}
	result := (data << 1) & 255
	result = result | carryIn
	carryOut := (data >> 7) == 1

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, carryOut)

	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.a.Write(result)
	}
}

func OperationRor(c *Cpu, addr *uint16) {
	c.clock.Tick()

	var data uint8
	if addr != nil {

		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {

		data = c.a.Read()
	}
	var carryIn uint8
	if c.status.Get(modules.CARRY) {
		carryIn = 1
	} else {
		carryIn = 0
	}
	result := data >> 1
	result = result | (carryIn << 7)
	carryOut := data == 1

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, carryOut)

	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.a.Write(result)
	}
}

func OperationBcc(c *Cpu, addr *uint16) {
	carryIn := c.status.Get(modules.CARRY)

	if !carryIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		// FIXEME the operater maybe reverse
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)

	}
}

func OperationBcs(c *Cpu, addr *uint16) {
	carryIn := c.status.Get(modules.CARRY)

	if carryIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//[relative] = BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)

	}
}

func OperationBeq(c *Cpu, addr *uint16) {

	zeroIn := c.status.Get(modules.ZERO)

	if zeroIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//[relative] = BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)
	}
}

func OperationBne(c *Cpu, addr *uint16) {

	zeroIn := c.status.Get(modules.ZERO)

	if !zeroIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//[relative] = BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)
	}
}

func OperationBvc(c *Cpu, addr *uint16) {

	overflowIn := c.status.Get(modules.OVERFLOW)

	if !overflowIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//[relative] = BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)
	}
}

func OperationBvs(c *Cpu, addr *uint16) {

	overflowIn := c.status.Get(modules.OVERFLOW)

	if overflowIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//[relative] = BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)
	}
}

func OperationBpl(c *Cpu, addr *uint16) {

	negativeIn := c.status.Get(modules.NEGATIVE)

	if !negativeIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//[relative] = BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)
	}
}

func OperationBmi(c *Cpu, addr *uint16) {

	negativeIn := c.status.Get(modules.NEGATIVE)

	if negativeIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(c.ram.Read(*addr))
		//[relative] = BitArray(uint=addr, length=8).unpack('int')
		pc := c.programCounter.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 < 256

		if carryOut {
			c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounter.Write(target)
	}
}

func OperationBit(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := a & memory

	overflow := memory>>6 == 1
	negative := memory>>7 == 1
	c.status.Change(modules.OVERFLOW, overflow)
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationJmp(c *Cpu, addr *uint16) {
	if false {
		// call yield due to be generator
		c.clock.Tick()
	}
	c.programCounter.Write(*addr - 1)
}

func OperationJsr(c *Cpu, addr *uint16) {
	c.clock.Tick()
	data := c.programCounter.Read()
	low := uint8(data & 255)
	high := uint8((data >> 8) & 255)
	c.clock.Tick()
	c.stack.Push(high)
	c.clock.Tick()
	c.stack.Push(low)
	target := *addr - 1
	c.programCounter.Write(target)
}

func OperationRts(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.programCounter.Increment()
	c.clock.Tick()
	c.clock.Tick()
	low_pc := c.stack.Pop()
	c.clock.Tick()
	high_pc := c.stack.Pop()
	c.clock.Tick()
	return_addr := (uint16(high_pc) << 8) + uint16(low_pc)
	c.programCounter.Write(return_addr)
}

func OperationBrk(c *Cpu, addr *uint16) {
	// TODO
	// it is wrong to check i flag here
	c.clock.Tick()

	interrupt := c.status.Get(modules.INTERRUPT)

	if !interrupt {
		c.status.Set(modules.BLEAK)

		next_pc := c.programCounter.Read() + 1

		low := uint8(next_pc & 255)
		high := uint8((next_pc >> 8) & 255)
		c.clock.Tick()
		c.stack.Push(high)
		c.clock.Tick()
		c.stack.Push(low)

		c.clock.Tick()
		c.stack.Push(c.status.Read())

		c.status.Set(modules.INTERRUPT)

		c.clock.Tick()
		low_pc := c.ram.Read(uint16(0xFFFE))
		c.clock.Tick()
		high_pc := c.ram.Read(uint16(0xFFFF))

		return_addr := (uint16(high_pc) << 8) + uint16(low_pc)
		c.programCounter.Write(return_addr)
	}
}

func OperationRti(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.programCounter.Increment()
	c.clock.Tick()

	c.clock.Tick()
	c.status.Write(uint8(c.stack.Pop()))
	c.clock.Tick()
	low_pc := c.stack.Pop()
	c.clock.Tick()
	high_pc := c.stack.Pop()
	return_addr := (uint16(high_pc) << 8) + uint16(low_pc)

	c.programCounter.Write(return_addr - 1)
}

func OperationCmp(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := a - memory
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, result >= 0)
}

func OperationCpx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.x.Read()

	memory := c.ram.Read(*addr)

	result := x - memory

	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, result >= 0)
}

func OperationCpy(c *Cpu, addr *uint16) {
	c.clock.Tick()

	y := c.y.Read()

	memory := c.ram.Read(*addr)

	result := y - memory

	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, result >= 0)
}

func OperationInc(c *Cpu, addr *uint16) {

	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.clock.Tick()
	result := (memory + 1) & 255
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
	c.clock.Tick()

	c.ram.Write(*addr, result)
}

func OperationDec(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.clock.Tick()
	result := (memory - 1) & 255
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)

	c.clock.Tick()

	c.ram.Write(*addr, result)
}

func OperationInx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.x.Read()

	result := (x + 1) & 255
	c.x.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationDex(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.x.Read()

	result := (x - 1) & 255
	c.x.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationIny(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.y.Read()
	result := (y + 1) & 255
	c.y.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationDey(c *Cpu, addr *uint16) {
	c.clock.Tick()

	y := c.y.Read()

	result := (y - 1) & 255
	c.y.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func OperationClc(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.CARRY)
}

func OperationSec(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Set(modules.CARRY)
}

func OperationCli(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.INTERRUPT)
}

func OperationSei(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Set(modules.INTERRUPT)
}

func OperationCld(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.DECIMAL)
}

func OperationSed(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Set(modules.DECIMAL)
}

func OperationClv(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.OVERFLOW)
}

func OperationLda(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.a.Write(memory)
	// TODO
	// status registers modified according to a register , not memory data
	negative := memory>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, memory == 0)
}

func OperationLdx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.x.Write(memory)

	// TODO
	// status registers modified according to x register , not memory data
	negative := memory>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, memory == 0)
}

func OperationLdy(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.y.Write(memory)

	// TODO
	// status registers modified according to y register , not memory data
	negative := memory>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, memory == 0)
}

func OperationSta(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	c.ram.Write(*addr, a)
}

func OperationStx(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.x.Read()

	c.ram.Write(*addr, x)
}

func OperationSty(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.y.Read()

	c.ram.Write(*addr, y)
}

func OperationTax(c *Cpu, addr *uint16) {
	c.clock.Tick()
	a := c.a.Read()
	c.x.Write(a)

	// TODO
	// status registers modified according to y register , not memory data
	negative := a>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, a == 0)
}

func OperationTxa(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.x.Read()
	c.a.Write(x)

	// TODO
	// status registers modified according to y register , not memory data
	negative := x>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, x == 0)
}

func OperationTay(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()
	c.y.Write(a)

	// TODO
	// status registers modified according to y register , not memory data
	negative := a>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, a == 0)
}

func OperationTya(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.y.Read()
	c.a.Write(y)

	// TODO
	// status registers modified according to y register , not memory data
	negative := y>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, y == 0)
}

func OperationTsx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	stackPointer := c.stack.GetStackPointer()
	c.x.Write(stackPointer)

	// TODO
	// status registers modified according to y register , not memory data
	negative := stackPointer>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, stackPointer == 0)
}

func OperationTxs(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.x.Read()
	c.stack.SetStackPointer(x)
}

func OperationPha(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.stack.Push(c.a.Read())
}

func OperationPla(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.clock.Tick()
	data := c.stack.Pop()
	c.a.Write(data)
	negative := data>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, data == 0)
}

func OperationPhp(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.stack.Push(c.status.Read())
}

func OperationPlp(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.clock.Tick()
	data := c.stack.Pop()
	c.status.Write(data)
}

func OperationNop(c *Cpu, addr *uint16) {
	c.clock.Tick()
}

// below is illegal opcode func definition
// combination of two operations with the same addring mode
func OperationSlo(c *Cpu, addr *uint16) {
	OperationAsl(c, addr)
	OperationOra(c, addr)
}

func OperationRla(c *Cpu, addr *uint16) {
	OperationRol(c, addr)
	OperationAnd(c, addr)
}

func OperationSre(c *Cpu, addr *uint16) {
	OperationLsr(c, addr)
	OperationEor(c, addr)
}

func OperationRra(c *Cpu, addr *uint16) {
	OperationRor(c, addr)
	OperationAdc(c, addr)
}

func OperationSax(c *Cpu, addr *uint16) {

	a := c.a.Read()
	x := c.x.Read()
	result := a & x
	c.clock.Tick()

	c.ram.Write(*addr, result)
}

func OperationLax(c *Cpu, addr *uint16) {
	OperationLda(c, addr)
	OperationLdx(c, addr)
}

func OperationDcp(c *Cpu, addr *uint16) {
	OperationDec(c, addr)
	OperationCmp(c, addr)
}

func OperationIsc(c *Cpu, addr *uint16) {
	OperationInc(c, addr)
	OperationSbc(c, addr)
}

// combinations of an immediate and an implied command
func OperationAnc(c *Cpu, addr *uint16) {
	OperationAnd(c, addr)

	a := c.a.Read()
	carryOut := (a >> 7) == 1

	c.status.Change(modules.CARRY, carryOut)
}

func OperationAlr(c *Cpu, addr *uint16) {
	OperationAnd(c, addr)
	OperationLsr(c, nil)
}

func OperationArr(c *Cpu, addr *uint16) {
	OperationAnd(c, addr)
	OperationRor(c, nil)
}

func OperationXaa(c *Cpu, addr *uint16) {
	OperationTxa(c, nil)
	OperationAnd(c, addr)
}

func OperationLaxi(c *Cpu, addr *uint16) {
	OperationLda(c, addr)
	OperationTax(c, nil)
}

func OperationAxs(c *Cpu, addr *uint16) {
	// nop
}

func OperationSbcn(c *Cpu, addr *uint16) {
	OperationSbc(c, addr)
	OperationNop(c, nil)
}

func OperationAhx(c *Cpu, addr *uint16) {}

func OperationShx(c *Cpu, addr *uint16) {}

func OperationShy(c *Cpu, addr *uint16) {}

func OperationTas(c *Cpu, addr *uint16) {}

func OperationLas(c *Cpu, addr *uint16) {}
