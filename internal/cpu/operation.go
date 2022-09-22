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
	Ora = &Operation{do: doOra, opecode: "ORA"}
	And = &Operation{do: doAnd, opecode: "AND"}
	Eor = &Operation{do: doEor, opecode: "EOR"}
	Adc = &Operation{do: doAdc, opecode: "ADC"}
	Sbc = &Operation{do: doSbc, opecode: "SBC"}
	Cmp = &Operation{do: doCmp, opecode: "CMP"}
	Cpx = &Operation{do: doCpx, opecode: "CPX"}
	Cpy = &Operation{do: doCpy, opecode: "CPY"}
	Dec = &Operation{do: doDec, opecode: "DEC"}
	Dex = &Operation{do: doDex, opecode: "DEX"}
	Dey = &Operation{do: doDey, opecode: "DEY"}
	Inc = &Operation{do: doInc, opecode: "INC"}
	Inx = &Operation{do: doInx, opecode: "INX"}
	Iny = &Operation{do: doIny, opecode: "INY"}
	Asl = &Operation{do: doAsl, opecode: "ASL"}
	Rol = &Operation{do: doRol, opecode: "ROL"}
	Lsr = &Operation{do: doLsr, opecode: "LSR"}
	Ror = &Operation{do: doRor, opecode: "ROR"}
	// move
	Lda = &Operation{do: doLda, opecode: "LDA"}
	Sta = &Operation{do: doSta, opecode: "STA"}
	Ldx = &Operation{do: doLdx, opecode: "LDX"}
	Stx = &Operation{do: doStx, opecode: "STX"}
	Ldy = &Operation{do: doLdy, opecode: "LDY"}
	Sty = &Operation{do: doSty, opecode: "STY"}
	Tax = &Operation{do: doTax, opecode: "TAX"}
	Txa = &Operation{do: doTxa, opecode: "TXA"}
	Tay = &Operation{do: doTay, opecode: "TAY"}
	Tya = &Operation{do: doTya, opecode: "TYA"}
	Tsx = &Operation{do: doTsx, opecode: "TSX"}
	Txs = &Operation{do: doTxs, opecode: "TXS"}
	Pla = &Operation{do: doPla, opecode: "PLA"}
	Pha = &Operation{do: doPha, opecode: "PHA"}
	Plp = &Operation{do: doPlp, opecode: "PLP"}
	Php = &Operation{do: doPhp, opecode: "PHP"}
	// jump
	Bpl = &Operation{do: doBpl, opecode: "BPL"}
	Bmi = &Operation{do: doBmi, opecode: "BMI"}
	Bvc = &Operation{do: doBvc, opecode: "BVC"}
	Bvs = &Operation{do: doBvs, opecode: "BVS"}
	Bcc = &Operation{do: doBcc, opecode: "BCC"}
	Bcs = &Operation{do: doBcs, opecode: "BCS"}
	Bne = &Operation{do: doBne, opecode: "BNE"}
	Beq = &Operation{do: doBeq, opecode: "BEQ"}
	Brk = &Operation{do: doBrk, opecode: "BRK"}
	Rti = &Operation{do: doRti, opecode: "RTI"}
	Jsr = &Operation{do: doJsr, opecode: "JSR"}
	Rts = &Operation{do: doRts, opecode: "RTS"}
	Jmp = &Operation{do: doJmp, opecode: "JMP"}
	Bit = &Operation{do: doBit, opecode: "BIT"}
	Clc = &Operation{do: doClc, opecode: "CLC"}
	Sec = &Operation{do: doSec, opecode: "SEC"}
	Cld = &Operation{do: doCld, opecode: "CLD"}
	Sed = &Operation{do: doSed, opecode: "SED"}
	Cli = &Operation{do: doCli, opecode: "CLI"}
	Sei = &Operation{do: doSei, opecode: "SEI"}
	Clv = &Operation{do: doClv, opecode: "CLV"}
	Nop = &Operation{do: doNop, opecode: "NOP"}
	// illegal
	Slo = &Operation{do: doSlo, opecode: "SLO"}
	Rla = &Operation{do: doRla, opecode: "RLA"}
	Sre = &Operation{do: doSre, opecode: "SRE"}
	Rra = &Operation{do: doRra, opecode: "RRA"}
	Sax = &Operation{do: doSax, opecode: "SAX"}
	Lax = &Operation{do: doLax, opecode: "LAX"}
	Dcp = &Operation{do: doDcp, opecode: "DCP"}
	Isc = &Operation{do: doIsc, opecode: "ISC"}
	Anc = &Operation{do: doAnc, opecode: "ANC"}
	Alr = &Operation{do: doAlr, opecode: "ALR"}
	Arr = &Operation{do: doArr, opecode: "ARR"}
	Xaa = &Operation{do: doXaa, opecode: "XAA"}
	Axs = &Operation{do: doAxs, opecode: "AXS"}
	Ahx = &Operation{do: doAhx, opecode: "AHX"}
	Shy = &Operation{do: doShy, opecode: "SHY"}
	Shx = &Operation{do: doShx, opecode: "SHX"}
	Tas = &Operation{do: doTas, opecode: "TAS"}
	Las = &Operation{do: doLas, opecode: "LAS"}

}

type Operation struct {
	do                 func(c *Cpu, addr *uint16)
	opecode            string
	isCombineoperation bool
}

type Operatable interface {
	Do(c *Cpu, addr *uint16)
	Opecode() string
	IsWriteOperation() bool
	IsCombineOperation() bool
	SetCombineOperation(v bool)
}

func (o *Operation) IsWriteOperation() bool {
	if o.opecode == "ASL" || o.opecode == "LSR" || o.opecode == "ROL" || o.opecode == "ROR" ||
		o.opecode == "INC" || o.opecode == "DEC" || o.opecode == "STA" || o.opecode == "STX" ||
		o.opecode == "STY" {
		return true
	}
	return false
}

func (o *Operation) SetCombineOperation(v bool) {
	o.isCombineoperation = v
}
func (o *Operation) IsCombineOperation() bool {
	if o.isCombineoperation {
		o.SetCombineOperation(false)
		return true
	}
	return false
}

func (o *Operation) Do(c *Cpu, addr *uint16) {
	o.do(c, addr)
}

func (o *Operation) Opecode() string {
	return o.opecode
}
func doAdc(c *Cpu, addr *uint16) {
	c.clock.Tick()
	a := c.aRegister.Read()
	memory := c.ram.Read(*addr)
	carryIn := c.statusRegister.Get(modules.CARRY)

	result, carryOut := modules.UnsignedAdd(a, memory, carryIn)
	_, overflow := modules.SignedAdd(int8(a), int8(memory), carryIn)
	c.aRegister.Write(result)
	c.statusRegister.Change(modules.OVERFLOW, overflow)
	c.statusRegister.Change(modules.NEGATIVE, result>>7 == 1)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, carryOut)
}

func doSbc(c *Cpu, addr *uint16) {
	c.clock.Tick()
	a := c.aRegister.Read()
	memory := c.ram.Read(*addr)
	carryIn := c.statusRegister.Get(modules.CARRY)

	result, carryOut := modules.UnsignedSub(a, memory, carryIn)
	_, overflow := modules.SignedSub(int8(a), int8(memory), carryIn)

	c.aRegister.Write(result)
	c.statusRegister.Change(modules.OVERFLOW, overflow)
	c.statusRegister.Change(modules.NEGATIVE, result>>7 == 1)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, carryOut)
}

func doAnd(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.aRegister.Read()

	memory := c.ram.Read(*addr)

	result := memory & a
	c.aRegister.Write(result)

	c.statusRegister.Change(modules.NEGATIVE, result > 127)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doOra(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.aRegister.Read()

	memory := c.ram.Read(*addr)

	result := memory | a
	c.aRegister.Write(result)

	c.statusRegister.Change(modules.NEGATIVE, result > 127)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doEor(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.aRegister.Read()

	memory := c.ram.Read(*addr)

	result := memory ^ a
	c.aRegister.Write(result)

	c.statusRegister.Change(modules.NEGATIVE, result > 127)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doAsl(c *Cpu, addr *uint16) {
	if !c.op.IsCombineOperation() {
		c.clock.Tick()
	}
	var data uint8
	if addr != nil {
		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {
		data = c.aRegister.Read()
	}

	result := data << 1 & 255
	carryOut := data>>7 == 1

	c.statusRegister.Change(modules.NEGATIVE, result > 127)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, carryOut)
	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.aRegister.Write(result)
	}
}

func doLsr(c *Cpu, addr *uint16) {
	if !c.op.IsCombineOperation() {
		c.clock.Tick()
	}
	var data uint8
	if addr != nil {
		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {
		data = c.aRegister.Read()
	}
	result := data >> 1
	carryOut := data&1 == 1

	c.statusRegister.Change(modules.NEGATIVE, result > 127)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, carryOut)
	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.aRegister.Write(result)
	}
}

func doRol(c *Cpu, addr *uint16) {
	if !c.op.IsCombineOperation() {
		c.clock.Tick()
	}

	var data uint8
	if addr != nil {
		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {

		data = c.aRegister.Read()
	}
	var carryIn uint8
	if c.statusRegister.Get(modules.CARRY) {
		carryIn = 1
	} else {
		carryIn = 0
	}
	result := data << 1 & 255
	result = result | carryIn
	carryOut := data>>7 == 1

	c.statusRegister.Change(modules.NEGATIVE, result > 127)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, carryOut)

	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.aRegister.Write(result)
	}
}

func doRor(c *Cpu, addr *uint16) {
	if !c.op.IsCombineOperation() {
		c.clock.Tick()
	}

	var data uint8
	if addr != nil {

		data = c.ram.Read(*addr)
		c.clock.Tick()
	} else {

		data = c.aRegister.Read()
	}
	var carryIn uint8
	if c.statusRegister.Get(modules.CARRY) {
		carryIn = 1
	} else {
		carryIn = 0
	}
	result := data >> 1
	result = result | (carryIn << 7)
	carryOut := data&1 == 1

	c.statusRegister.Change(modules.NEGATIVE, result > 127)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, carryOut)

	if addr != nil {
		c.clock.Tick()
		c.ram.Write(*addr, result)
	} else {
		c.aRegister.Write(result)
	}
}

func doBcc(c *Cpu, addr *uint16) {
	carryIn := c.statusRegister.Get(modules.CARRY)
	if !carryIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		// FIXEME the operater maybe reverse
		carryOut := pc_low_int16 > 255

		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)

	}
}

func doBcs(c *Cpu, addr *uint16) {
	carryIn := c.statusRegister.Get(modules.CARRY)

	if carryIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 > 255

		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)

	}
}

func doBeq(c *Cpu, addr *uint16) {

	zeroIn := c.statusRegister.Get(modules.ZERO)

	if zeroIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 > 255
		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)
	}
}

func doBne(c *Cpu, addr *uint16) {

	zeroIn := c.statusRegister.Get(modules.ZERO)

	if !zeroIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 > 255

		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)
	}
}

func doBvc(c *Cpu, addr *uint16) {

	overflowIn := c.statusRegister.Get(modules.OVERFLOW)

	if !overflowIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 > 255

		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)
	}
}

func doBvs(c *Cpu, addr *uint16) {

	overflowIn := c.statusRegister.Get(modules.OVERFLOW)
	if overflowIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 > 255

		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)
	}
}

func doBpl(c *Cpu, addr *uint16) {

	negativeIn := c.statusRegister.Get(modules.NEGATIVE)

	if !negativeIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 > 255

		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)
	}
}

func doBmi(c *Cpu, addr *uint16) {

	negativeIn := c.statusRegister.Get(modules.NEGATIVE)

	if negativeIn {
		c.clock.Tick()
		// maybe + 1 is required for pc
		relative := int8(*addr)
		pc := c.programCounterRegister.Read()
		pc_low := pc & 255
		pc_high := (pc >> 8) & 255
		pc_low_int16 := int16(pc_low) + int16(relative)
		carryOut := pc_low_int16 > 255

		if carryOut {
			//c.clock.Tick()
		}
		target := (pc_high << 8) + uint16(pc_low_int16)
		c.programCounterRegister.Write(target)
	}
}

func doBit(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.aRegister.Read()

	memory := c.ram.Read(*addr)

	result := a & memory

	overflow := memory>>6&1 == 1
	negative := memory>>7 == 1
	c.statusRegister.Change(modules.OVERFLOW, overflow)
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doJmp(c *Cpu, addr *uint16) {
	if false {
		// call yield due to be generator
		c.clock.Tick()
	}
	c.programCounterRegister.Write(*addr - 1)
}

func doJsr(c *Cpu, addr *uint16) {
	c.clock.Tick()
	data := c.programCounterRegister.Read()
	low := uint8(data & 255)
	high := uint8((data >> 8) & 255)
	c.clock.Tick()
	c.stack.Push(high)
	c.clock.Tick()
	c.stack.Push(low)
	target := *addr - 1
	c.programCounterRegister.Write(target)
}

func doRts(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	c.clock.Tick()
	c.clock.Tick()
	low_pc := c.stack.Pop()
	c.clock.Tick()
	high_pc := c.stack.Pop()
	c.clock.Tick()
	return_addr := (uint16(high_pc) << 8) + uint16(low_pc)
	c.programCounterRegister.Write(return_addr)
}

func doBrk(c *Cpu, addr *uint16) {
	// TODO
	// it is wrong to check i flag here
	c.clock.Tick()

	interrupt := c.statusRegister.Get(modules.INTERRUPT)

	if !interrupt {
		c.statusRegister.Set(modules.BLEAK)

		next_pc := c.programCounterRegister.Read() + 1

		low := uint8(next_pc & 255)
		high := uint8((next_pc >> 8) & 255)
		c.clock.Tick()
		c.stack.Push(high)
		c.clock.Tick()
		c.stack.Push(low)

		c.clock.Tick()
		c.stack.Push(c.statusRegister.Read())

		c.statusRegister.Set(modules.INTERRUPT)

		c.clock.Tick()
		low_pc := c.ram.Read(uint16(0xFFFE))
		c.clock.Tick()
		high_pc := c.ram.Read(uint16(0xFFFF))

		return_addr := (uint16(high_pc) << 8) + uint16(low_pc)
		c.programCounterRegister.Write(return_addr)
	}
}

func doRti(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	c.clock.Tick()

	c.clock.Tick()
	c.statusRegister.Write(uint8(c.stack.Pop()))
	c.clock.Tick()
	low_pc := c.stack.Pop()
	c.clock.Tick()
	high_pc := c.stack.Pop()
	return_addr := (uint16(high_pc) << 8) + uint16(low_pc)

	c.programCounterRegister.Write(return_addr - 1)
}

func doCmp(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.aRegister.Read()

	memory := c.ram.Read(*addr)

	result := a - memory
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, a >= memory)
}

func doCpx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.xRegister.Read()

	memory := c.ram.Read(*addr)

	result := x - memory
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, x >= memory)
}

func doCpy(c *Cpu, addr *uint16) {
	c.clock.Tick()

	y := c.yRegister.Read()

	memory := c.ram.Read(*addr)

	result := y - memory

	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.statusRegister.Change(modules.CARRY, y >= memory)
}

func doInc(c *Cpu, addr *uint16) {
	if !c.op.IsCombineOperation() {
		c.clock.Tick()

	}

	memory := c.ram.Read(*addr)

	c.clock.Tick()
	result := (memory + 1) & 255
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
	c.clock.Tick()

	c.ram.Write(*addr, result)
}

func doDec(c *Cpu, addr *uint16) {
	if !c.op.IsCombineOperation() {
		c.clock.Tick()
	}

	memory := c.ram.Read(*addr)

	c.clock.Tick()
	result := (memory - 1) & 255
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)

	c.clock.Tick()

	c.ram.Write(*addr, result)
}

func doInx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.xRegister.Read()

	result := (x + 1) & 255
	c.xRegister.Write(result)
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doDex(c *Cpu, addr *uint16) {
	c.clock.Tick()

	x := c.xRegister.Read()

	result := (x - 1) & 255
	c.xRegister.Write(result)
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doIny(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.yRegister.Read()
	result := (y + 1) & 255
	c.yRegister.Write(result)
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doDey(c *Cpu, addr *uint16) {
	c.clock.Tick()

	y := c.yRegister.Read()

	result := (y - 1) & 255
	c.yRegister.Write(result)
	negative := result>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, result == 0)
}

func doClc(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.statusRegister.Clear(modules.CARRY)
}

func doSec(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.statusRegister.Set(modules.CARRY)
}

func doCli(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.statusRegister.Clear(modules.INTERRUPT)
}

func doSei(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.statusRegister.Set(modules.INTERRUPT)
}

func doCld(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.statusRegister.Clear(modules.DECIMAL)
}

func doSed(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.statusRegister.Set(modules.DECIMAL)
}

func doClv(c *Cpu, addr *uint16) {
	c.clock.Tick()

	c.statusRegister.Clear(modules.OVERFLOW)
}

func doLda(c *Cpu, addr *uint16) {
	if !c.op.IsCombineOperation() {
		c.clock.Tick()
	}

	memory := c.ram.Read(*addr)
	c.aRegister.Write(memory)
	negative := memory>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, memory == 0)
}

func doLdx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.xRegister.Write(memory)
	negative := memory>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, memory == 0)
}

func doLdy(c *Cpu, addr *uint16) {
	c.clock.Tick()

	memory := c.ram.Read(*addr)

	c.yRegister.Write(memory)
	negative := memory>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, memory == 0)
}

func doSta(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.aRegister.Read()

	c.ram.Write(*addr, a)
}

func doStx(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.xRegister.Read()

	c.ram.Write(*addr, x)
}

func doSty(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.yRegister.Read()

	c.ram.Write(*addr, y)

}

func doTax(c *Cpu, addr *uint16) {
	c.clock.Tick()
	a := c.aRegister.Read()
	c.xRegister.Write(a)
	negative := a>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, a == 0)
}

func doTxa(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.xRegister.Read()
	c.aRegister.Write(x)
	negative := x>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, x == 0)
}

func doTay(c *Cpu, addr *uint16) {
	c.clock.Tick()

	a := c.aRegister.Read()
	c.yRegister.Write(a)
	negative := a>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, a == 0)
}

func doTya(c *Cpu, addr *uint16) {
	c.clock.Tick()
	y := c.yRegister.Read()
	c.aRegister.Write(y)
	negative := y>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, y == 0)
}

func doTsx(c *Cpu, addr *uint16) {
	c.clock.Tick()

	stackPointer := c.stack.GetStackPointer()
	c.xRegister.Write(stackPointer)
	negative := stackPointer>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, stackPointer == 0)
}

func doTxs(c *Cpu, addr *uint16) {
	c.clock.Tick()
	x := c.xRegister.Read()
	c.stack.SetStackPointer(x)
}

func doPha(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.stack.Push(c.aRegister.Read())
}

func doPla(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.clock.Tick()
	data := c.stack.Pop()
	c.aRegister.Write(data)
	negative := data>>7 == 1
	c.statusRegister.Change(modules.NEGATIVE, negative)
	c.statusRegister.Change(modules.ZERO, data == 0)
}

func doPhp(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.stack.Push(c.statusRegister.Read())
}

func doPlp(c *Cpu, addr *uint16) {
	c.clock.Tick()
	c.clock.Tick()

	c.clock.Tick()
	data := c.stack.Pop()
	c.statusRegister.Write(data)
}

func doNop(c *Cpu, addr *uint16) {
	c.clock.Tick()
}

// below is illegal opcode func definition
// combination of two operations with the same addring mode
func doSlo(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doAsl(c, addr)
	doOra(c, addr)
}

func doRla(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doRol(c, addr)
	doAnd(c, addr)
}

func doSre(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doLsr(c, addr)
	doEor(c, addr)
}

func doRra(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doRor(c, addr)
	doAdc(c, addr)
}

func doSax(c *Cpu, addr *uint16) {

	a := c.aRegister.Read()
	x := c.xRegister.Read()
	result := a & x
	c.clock.Tick()

	c.ram.Write(*addr, result)
}

func doLax(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doLda(c, addr)
	doLdx(c, addr)
}

func doDcp(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doDec(c, addr)
	doCmp(c, addr)
}

func doIsc(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doInc(c, addr)
	doSbc(c, addr)
}

// combinations of an immediate and an implied command
func doAnc(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doAnd(c, addr)

	a := c.aRegister.Read()
	carryOut := a>>7 == 1

	c.statusRegister.Change(modules.CARRY, carryOut)
}

func doAlr(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doAnd(c, addr)
	doLsr(c, nil)
}

func doArr(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doAnd(c, addr)
	doRor(c, nil)
}

func doXaa(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doTxa(c, nil)
	doAnd(c, addr)
}

func doLaxi(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doLda(c, addr)
	doTax(c, nil)
}

func doAxs(c *Cpu, addr *uint16) {
	// nop
}

func doSbcn(c *Cpu, addr *uint16) {
	c.op.SetCombineOperation(true)
	doSbc(c, addr)
	doNop(c, nil)
}

func doAhx(c *Cpu, addr *uint16) {}

func doShx(c *Cpu, addr *uint16) {}

func doShy(c *Cpu, addr *uint16) {}

func doTas(c *Cpu, addr *uint16) {}

func doLas(c *Cpu, addr *uint16) {}
