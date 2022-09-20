package cpu

import "github.com/u400r/gones/internal/modules"

var (
	// logical and alithmetic
	Ora, And, Eor, Adc, Sbc, Cmp, Cpx, Cpy, Dec, Dex, Dey, Inc, Inx, Iny, Asl, Rol, Lsr, Ror *Operation
	// move
	Lda, Sta, Ldx, Stx, Ldy, Sty, Tax, Txa, Tay, Tya, Tsx, Txs, Pla, Pha, Plp, Php *Operation
	// jump
	Bpl, Bmi, Bvc, Bvs, Bcc, Bcs, Bne, Beq, Brk, Rti, Jsr, Rts, Jmp, Bit, Clc, Sec, Cld, Sed, Cli, Sei, Clv, Nop *Operation
	// illegal
	Slo, Rla, Sre, Rra, Sax, Lax, Dcp, Isc, Anc, Alr, Arr, Xaa, Axs, Ahx, Shy, Shx, Tas, Las *Operation
)

func init() {
	Ora = &Operation{do: doOra, opecode: "Ora"}
	And = &Operation{do: doAnd, opecode: "And"}
	Eor = &Operation{do: doEor, opecode: "Eor"}
	Adc = &Operation{do: doAdc, opecode: "Adc"}
	Sbc = &Operation{do: doSbc, opecode: "Sbc"}
	Cmp = &Operation{do: doCmp, opecode: "Cmp"}
	Cpx = &Operation{do: doCpx, opecode: "Cpx"}
	Cpy = &Operation{do: doCpy, opecode: "Cpy"}
	Dec = &Operation{do: doDec, opecode: "Dec"}
	Dex = &Operation{do: doDex, opecode: "Dex"}
	Dey = &Operation{do: doDey, opecode: "Dey"}
	Inc = &Operation{do: doInc, opecode: "Inc"}
	Inx = &Operation{do: doInx, opecode: "Inx"}
	Iny = &Operation{do: doIny, opecode: "Iny"}
	Asl = &Operation{do: doAsl, opecode: "Asl"}
	Rol = &Operation{do: doRol, opecode: "Rol"}
	Lsr = &Operation{do: doLsr, opecode: "Lsr"}
	Ror = &Operation{do: doRor, opecode: "Ror"}
	// move
	Lda = &Operation{do: doLda, opecode: "Lda"}
	Sta = &Operation{do: doSta, opecode: "Sta"}
	Ldx = &Operation{do: doLdx, opecode: "Ldx"}
	Stx = &Operation{do: doStx, opecode: "Stx"}
	Ldy = &Operation{do: doLdy, opecode: "Ldy"}
	Sty = &Operation{do: doSty, opecode: "Sty"}
	Tax = &Operation{do: doTax, opecode: "Tax"}
	Txa = &Operation{do: doTxa, opecode: "Txa"}
	Tay = &Operation{do: doTay, opecode: "Tay"}
	Tya = &Operation{do: doTya, opecode: "Tya"}
	Tsx = &Operation{do: doTsx, opecode: "Tsx"}
	Txs = &Operation{do: doTxs, opecode: "Txs"}
	Pla = &Operation{do: doPla, opecode: "Pla"}
	Pha = &Operation{do: doPha, opecode: "Pha"}
	Plp = &Operation{do: doPlp, opecode: "Plp"}
	Php = &Operation{do: doPhp, opecode: "Php"}
	// jump
	Bpl = &Operation{do: doBpl, opecode: "Bpl"}
	Bmi = &Operation{do: doBmi, opecode: "Bmi"}
	Bvc = &Operation{do: doBvc, opecode: "Bvc"}
	Bvs = &Operation{do: doBvs, opecode: "Bvs"}
	Bcc = &Operation{do: doBcc, opecode: "Bcc"}
	Bcs = &Operation{do: doBcs, opecode: "Bcs"}
	Bne = &Operation{do: doBne, opecode: "Bne"}
	Beq = &Operation{do: doBeq, opecode: "Beq"}
	Brk = &Operation{do: doBrk, opecode: "Brk"}
	Rti = &Operation{do: doRti, opecode: "Rti"}
	Jsr = &Operation{do: doJsr, opecode: "Jsr"}
	Rts = &Operation{do: doRts, opecode: "Rts"}
	Jmp = &Operation{do: doJmp, opecode: "Jmp"}
	Bit = &Operation{do: doBit, opecode: "Bit"}
	Clc = &Operation{do: doClc, opecode: "Clc"}
	Sec = &Operation{do: doSec, opecode: "Sec"}
	Cld = &Operation{do: doCld, opecode: "Cld"}
	Sed = &Operation{do: doSed, opecode: "Sed"}
	Cli = &Operation{do: doCli, opecode: "Cli"}
	Sei = &Operation{do: doSei, opecode: "Sei"}
	Clv = &Operation{do: doClv, opecode: "Clv"}
	Nop = &Operation{do: doNop, opecode: "Nop"}
	// illegal
	Slo = &Operation{do: doSlo, opecode: "Slo"}
	Rla = &Operation{do: doRla, opecode: "Rla"}
	Sre = &Operation{do: doSre, opecode: "Sre"}
	Rra = &Operation{do: doRra, opecode: "Rra"}
	Sax = &Operation{do: doSax, opecode: "Sax"}
	Lax = &Operation{do: doLax, opecode: "Lax"}
	Dcp = &Operation{do: doDcp, opecode: "Dcp"}
	Isc = &Operation{do: doIsc, opecode: "Isc"}
	Anc = &Operation{do: doAnc, opecode: "Anc"}
	Alr = &Operation{do: doAlr, opecode: "Alr"}
	Arr = &Operation{do: doArr, opecode: "Arr"}
	Xaa = &Operation{do: doXaa, opecode: "Xaa"}
	Axs = &Operation{do: doAxs, opecode: "Axs"}
	Ahx = &Operation{do: doAhx, opecode: "Ahx"}
	Shy = &Operation{do: doShy, opecode: "Shy"}
	Shx = &Operation{do: doShx, opecode: "Shx"}
	Tas = &Operation{do: doTas, opecode: "Tas"}
	Las = &Operation{do: doLas, opecode: "Las"}

}

type Operation struct {
	do      func(c *Cpu, addr *uint16)
	opecode string
}

type Operatable interface {
	Do(c *Cpu, addr *uint16)
	Opecode() string
}

func (o *Operation) Do(c *Cpu, addr *uint16) {
	o.do(c, addr)
}

func (o *Operation) Opecode() string {
	return o.opecode
}
func doAdc(c *Cpu, addr *uint16) {
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

func doSbc(c *Cpu, addr *uint16) {
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

func doAnd(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := memory & a
	c.a.Write(result)

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
}

func doOra(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := memory | a
	c.a.Write(result)

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
}

func doEor(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := memory ^ a
	c.a.Write(result)

	c.status.Change(modules.NEGATIVE, result > 127)
	c.status.Change(modules.ZERO, result == 0)
}

func doAsl(c *Cpu, addr *uint16) {
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

func doLsr(c *Cpu, addr *uint16) {
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

func doRol(c *Cpu, addr *uint16) {
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

func doRor(c *Cpu, addr *uint16) {
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

func doBcc(c *Cpu, addr *uint16) {
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

func doBcs(c *Cpu, addr *uint16) {
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

func doBeq(c *Cpu, addr *uint16) {

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

func doBne(c *Cpu, addr *uint16) {

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

func doBvc(c *Cpu, addr *uint16) {

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

func doBvs(c *Cpu, addr *uint16) {

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

func doBpl(c *Cpu, addr *uint16) {

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

func doBmi(c *Cpu, addr *uint16) {

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

func doBit(c *Cpu, addr *uint16) {
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

func doJmp(c *Cpu, addr *uint16) {
	if false {
		// call yield due to be generator
		c.clock.Tick()
	}
	c.programCounter.Write(*addr - 1)
}

func doJsr(c *Cpu, addr *uint16) {
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

func doRts(c *Cpu, addr *uint16) {
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

func doBrk(c *Cpu, addr *uint16) {
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

func doRti(c *Cpu, addr *uint16) {
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

func doCmp(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	memory := c.ram.Read(*addr)

	result := a - memory
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, result >= 0)
}

func doCpx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.x.Read()

	memory := c.ram.Read(*addr)

	result := x - memory

	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, result >= 0)
}

func doCpy(c *Cpu, addr *uint16) {
	c.clock.Tick()

	y := c.y.Read()

	memory := c.ram.Read(*addr)

	result := y - memory

	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
	c.status.Change(modules.CARRY, result >= 0)
}

func doInc(c *Cpu, addr *uint16) {

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

func doDec(c *Cpu, addr *uint16) {
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

func doInx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.x.Read()

	result := (x + 1) & 255
	c.x.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func doDex(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.x.Read()

	result := (x - 1) & 255
	c.x.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func doIny(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.y.Read()
	result := (y + 1) & 255
	c.y.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func doDey(c *Cpu, addr *uint16) {
	c.clock.Tick()

	y := c.y.Read()

	result := (y - 1) & 255
	c.y.Write(result)
	negative := result>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, result == 0)
}

func doClc(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.CARRY)
}

func doSec(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Set(modules.CARRY)
}

func doCli(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.INTERRUPT)
}

func doSei(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Set(modules.INTERRUPT)
}

func doCld(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.DECIMAL)
}

func doSed(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Set(modules.DECIMAL)
}

func doClv(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.status.Clear(modules.OVERFLOW)
}

func doLda(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.a.Write(memory)
	// TODO
	// status registers modified according to a register , not memory data
	negative := memory>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, memory == 0)
}

func doLdx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.x.Write(memory)

	// TODO
	// status registers modified according to x register , not memory data
	negative := memory>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, memory == 0)
}

func doLdy(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.y.Write(memory)

	// TODO
	// status registers modified according to y register , not memory data
	negative := memory>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, memory == 0)
}

func doSta(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()

	c.ram.Write(*addr, a)
}

func doStx(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.x.Read()

	c.ram.Write(*addr, x)
}

func doSty(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.y.Read()

	c.ram.Write(*addr, y)
}

func doTax(c *Cpu, addr *uint16) {
	c.clock.Tick()
	a := c.a.Read()
	c.x.Write(a)

	// TODO
	// status registers modified according to y register , not memory data
	negative := a>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, a == 0)
}

func doTxa(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.x.Read()
	c.a.Write(x)

	// TODO
	// status registers modified according to y register , not memory data
	negative := x>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, x == 0)
}

func doTay(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.a.Read()
	c.y.Write(a)

	// TODO
	// status registers modified according to y register , not memory data
	negative := a>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, a == 0)
}

func doTya(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.y.Read()
	c.a.Write(y)

	// TODO
	// status registers modified according to y register , not memory data
	negative := y>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, y == 0)
}

func doTsx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	stackPointer := c.stack.GetStackPointer()
	c.x.Write(stackPointer)

	// TODO
	// status registers modified according to y register , not memory data
	negative := stackPointer>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, stackPointer == 0)
}

func doTxs(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.x.Read()
	c.stack.SetStackPointer(x)
}

func doPha(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.stack.Push(c.a.Read())
}

func doPla(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.clock.Tick()
	data := c.stack.Pop()
	c.a.Write(data)
	negative := data>>7 == 1
	c.status.Change(modules.NEGATIVE, negative)
	c.status.Change(modules.ZERO, data == 0)
}

func doPhp(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.stack.Push(c.status.Read())
}

func doPlp(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.clock.Tick()
	data := c.stack.Pop()
	c.status.Write(data)
}

func doNop(c *Cpu, addr *uint16) {
	c.clock.Tick()
}

// below is illegal opcode func definition
// combination of two operations with the same addring mode
func doSlo(c *Cpu, addr *uint16) {
	doAsl(c, addr)
	doOra(c, addr)
}

func doRla(c *Cpu, addr *uint16) {
	doRol(c, addr)
	doAnd(c, addr)
}

func doSre(c *Cpu, addr *uint16) {
	doLsr(c, addr)
	doEor(c, addr)
}

func doRra(c *Cpu, addr *uint16) {
	doRor(c, addr)
	doAdc(c, addr)
}

func doSax(c *Cpu, addr *uint16) {

	a := c.a.Read()
	x := c.x.Read()
	result := a & x
	c.clock.Tick()

	c.ram.Write(*addr, result)
}

func doLax(c *Cpu, addr *uint16) {
	doLda(c, addr)
	doLdx(c, addr)
}

func doDcp(c *Cpu, addr *uint16) {
	doDec(c, addr)
	doCmp(c, addr)
}

func doIsc(c *Cpu, addr *uint16) {
	doInc(c, addr)
	doSbc(c, addr)
}

// combinations of an immediate and an implied command
func doAnc(c *Cpu, addr *uint16) {
	doAnd(c, addr)

	a := c.a.Read()
	carryOut := (a >> 7) == 1

	c.status.Change(modules.CARRY, carryOut)
}

func doAlr(c *Cpu, addr *uint16) {
	doAnd(c, addr)
	doLsr(c, nil)
}

func doArr(c *Cpu, addr *uint16) {
	doAnd(c, addr)
	doRor(c, nil)
}

func doXaa(c *Cpu, addr *uint16) {
	doTxa(c, nil)
	doAnd(c, addr)
}

func doLaxi(c *Cpu, addr *uint16) {
	doLda(c, addr)
	doTax(c, nil)
}

func doAxs(c *Cpu, addr *uint16) {
	// nop
}

func doSbcn(c *Cpu, addr *uint16) {
	doSbc(c, addr)
	doNop(c, nil)
}

func doAhx(c *Cpu, addr *uint16) {}

func doShx(c *Cpu, addr *uint16) {}

func doShy(c *Cpu, addr *uint16) {}

func doTas(c *Cpu, addr *uint16) {}

func doLas(c *Cpu, addr *uint16) {}
