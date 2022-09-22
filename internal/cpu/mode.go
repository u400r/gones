package cpu

var (
	Immediate, Absolute, AbsoluteX, AbsoluteY Addressing
	ZeroPage, ZeroPageX, ZeroPageY, Relative  Addressing
	Indirect, IndirectX, IndirectY, Implicit  Addressing
)

func init() {
	Immediate = &ImmediateMode{}
	Absolute = &AbsoluteMode{}
	AbsoluteX = &AbsoluteXMode{}
	AbsoluteY = &AbsoluteYMode{}
	ZeroPage = &ZeroPageMode{}
	ZeroPageX = &ZeroPageXMode{}
	ZeroPageY = &ZeroPageYMode{}
	Relative = &RelativeMode{}
	Indirect = &IndirectMode{}
	IndirectX = &IndirectXMode{}
	IndirectY = &IndirectYMode{}
	Implicit = &ImplicitMode{}
}

type Addressing interface {
	GetAddress(c *Cpu) *uint16
	GetModeString() string
	GetOperandBytes(c *Cpu) (*byte, *byte)
	isWaitOneClock(c *Cpu) bool
	isNop(c *Cpu) bool
}

type Mode struct {
	firstByte  *uint8
	secondByte *uint8
}

func (m *Mode) isWaitOneClock(c *Cpu) bool {
	if c.op.IsWriteOperation() {
		return true
	}
	return false
}

func (m *Mode) isNop(c *Cpu) bool {
	if c.op.Opecode() == "NOP" {
		return true
	} else {
		return false
	}
}

type ImmediateMode struct {
	Mode
}

func (m *Mode) GetOperandBytes(c *Cpu) (*byte, *byte) {
	return m.firstByte, m.secondByte
}

func (i *ImmediateMode) GetAddress(c *Cpu) *uint16 {
	if false {
		c.clock.Tick()
	}
	c.programCounterRegister.Increment()
	address := c.programCounterRegister.Read()
	return &address

}
func (i *ImmediateMode) GetModeString() string {
	return "Immediate"
}

type AbsoluteMode struct {
	Mode
}

func (a *AbsoluteMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	absolute_low := c.ram.Read(c.programCounterRegister.Read())
	a.firstByte = &absolute_low
	c.clock.Tick()
	c.programCounterRegister.Increment()
	absolute_high := c.ram.Read(c.programCounterRegister.Read())
	a.secondByte = &absolute_high
	address := uint16(absolute_high)<<8 + uint16(absolute_low)
	return &address

}

func (a *AbsoluteMode) GetModeString() string {
	return "Absolute"
}

type AbsoluteXMode struct {
	Mode
}

func (a *AbsoluteXMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	absolute_low := c.ram.Read(c.programCounterRegister.Read())
	a.firstByte = &absolute_low
	c.clock.Tick()
	c.programCounterRegister.Increment()
	absolute_high := c.ram.Read(c.programCounterRegister.Read())
	a.secondByte = &absolute_high
	index := c.xRegister.Read()
	// why this calculation does not take 1 clock?
	absolute_low_16 := uint16(absolute_low) + uint16(index)
	carry_out := absolute_low_16 > 255
	if carry_out {
		// tick one more clock but no need to Increment
		// absolute_high because carry is included in low
		c.clock.Tick()
	} else {
		if a.isWaitOneClock(c) {
			c.clock.Tick()
		}

	}
	address := uint16(absolute_high)<<8 + absolute_low_16
	return &address

}

func (a *AbsoluteXMode) GetModeString() string {
	return "AbsoluteX"
}

type AbsoluteYMode struct {
	Mode
}

func (a *AbsoluteYMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	absolute_low := c.ram.Read(c.programCounterRegister.Read())
	a.firstByte = &absolute_low
	c.clock.Tick()
	c.programCounterRegister.Increment()
	absolute_high := c.ram.Read(c.programCounterRegister.Read())
	a.secondByte = &absolute_high
	index := c.yRegister.Read()
	// why this calculation does not take 1 clock?
	absolute_low_16 := uint16(absolute_low) + uint16(index)
	carry_out := absolute_low_16 > 255
	if carry_out {
		// tick one more clock but no need to Increment
		// absolute_high because carry is included in low
		c.clock.Tick()
	} else {
		if a.isWaitOneClock(c) {
			c.clock.Tick()
		}
	}

	address := uint16(absolute_high)<<8 + absolute_low_16
	return &address

}
func (a *AbsoluteYMode) GetModeString() string {
	return "AbsoluteY"
}

type ZeroPageMode struct {
	Mode
}

func (z *ZeroPageMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	address_low := c.ram.Read(c.programCounterRegister.Read())
	z.firstByte = &address_low
	address := uint16(address_low)
	return &address

}

func (a *ZeroPageMode) GetModeString() string {
	return "ZeroPage"
}

type ZeroPageXMode struct {
	Mode
}

func (z *ZeroPageXMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	zero_page_address := c.ram.Read(c.programCounterRegister.Read())
	z.firstByte = &zero_page_address
	c.clock.Tick()
	// ignore carry
	zero_page_x_address := uint16((zero_page_address + c.xRegister.Read()) & 255)
	return &zero_page_x_address
}
func (a *ZeroPageXMode) GetModeString() string {
	return "ZeroPageX"
}

type ZeroPageYMode struct {
	Mode
}

func (z *ZeroPageYMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	zero_page_address := c.ram.Read(c.programCounterRegister.Read())
	z.firstByte = &zero_page_address
	c.clock.Tick()
	zero_page_y_address := uint16((zero_page_address + c.yRegister.Read()) & 255)
	return &zero_page_y_address

}
func (a *ZeroPageYMode) GetModeString() string {
	return "ZeroPageY"
}

type RelativeMode struct {
	Mode
}

func (r *RelativeMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	relative := c.ram.Read(c.programCounterRegister.Read())
	r.firstByte = &relative
	relative_16 := uint16(relative)
	return &relative_16

}
func (r *RelativeMode) GetModeString() string {
	return "Relative"
}

type IndirectMode struct {
	Mode
}

func (i *IndirectMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	indirect_address_low := c.ram.Read(c.programCounterRegister.Read())
	i.firstByte = &indirect_address_low
	c.clock.Tick()
	c.programCounterRegister.Increment()
	indirect_address_high := c.ram.Read(c.programCounterRegister.Read())
	i.secondByte = &indirect_address_high
	c.clock.Tick()
	address_low := c.ram.Read(uint16(indirect_address_high)<<8 + uint16(indirect_address_low))
	c.clock.Tick()
	// ignore carry from lower bit inrement
	address_high := c.ram.Read(uint16(indirect_address_high)<<8 + uint16(indirect_address_low+1))
	address := uint16(address_high)<<8 + uint16(address_low)
	return &address

}
func (i *IndirectMode) GetModeString() string {
	return "Indirect"
}

type IndirectXMode struct {
	Mode
}

func (i *IndirectXMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	zero_page_address := c.ram.Read(c.programCounterRegister.Read())
	i.firstByte = &zero_page_address
	if i.isNop(c) {
		return nil
	}
	// ignore carry_out so this shall lower than 0xff
	c.clock.Tick()
	indirect_address := (zero_page_address + c.xRegister.Read()) & 0xff
	c.clock.Tick()
	address_low := c.ram.Read(uint16(indirect_address))
	c.clock.Tick()
	address_high := c.ram.Read(uint16((indirect_address + 1) & 0xff))
	address := uint16(address_high)<<8 + uint16(address_low)
	return &address

}

func (i *IndirectXMode) GetModeString() string {
	return "IndirectX"
}

type IndirectYMode struct {
	Mode
}

func (i *IndirectYMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounterRegister.Increment()
	indirect_address := c.ram.Read(c.programCounterRegister.Read())
	i.firstByte = &indirect_address
	c.clock.Tick()
	address_low := c.ram.Read(uint16(indirect_address))
	c.clock.Tick()
	address_high := c.ram.Read(uint16(indirect_address + 1))
	index := c.yRegister.Read()
	address_low_16 := uint16(address_low) + uint16(index)
	carry_out := address_low_16 > 255
	if carry_out {
		c.clock.Tick()
	} else {
		if i.isWaitOneClock(c) {
			c.clock.Tick()
		}
	}
	address := uint16(address_high)<<8 + address_low_16
	return &address

}

func (i *IndirectYMode) GetModeString() string {
	return "IndirectY"
}

type ImplicitMode struct {
	Mode
}

func (i *ImplicitMode) GetAddress(c *Cpu) *uint16 {
	if false {
		c.clock.Tick()
	}
	return nil
}

func (i *ImplicitMode) GetModeString() string {
	return "Implicit"
}
