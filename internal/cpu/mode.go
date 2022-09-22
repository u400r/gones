package cpu

var (
	Immediate, Abusolute, AbusoluteX, AbusoluteY Addressing
	ZeroPage, ZeroPageX, ZeroPageY, Relative     Addressing
	Indirect, IndirectX, IndirectY, Implicit     Addressing
)

func init() {
	Immediate = &ImmediateMode{}
	Abusolute = &AbusoluteMode{}
	AbusoluteX = &AbusoluteXMode{}
	AbusoluteY = &AbusoluteYMode{}
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
}

type Mode struct {
	firstByte  *uint8
	secondByte *uint8
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
	c.programCounter.Increment()
	address := c.programCounter.Read()
	return &address

}
func (i *ImmediateMode) GetModeString() string {
	return "Immediate"
}

type AbusoluteMode struct {
	Mode
}

func (a *AbusoluteMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_low := c.ram.Read(c.programCounter.Read())
	a.firstByte = &absolute_low
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_high := c.ram.Read(c.programCounter.Read())
	a.secondByte = &absolute_high
	address := uint16(absolute_high)<<8 + uint16(absolute_low)
	return &address

}

func (a *AbusoluteMode) GetModeString() string {
	return "Abusolute"
}

type AbusoluteXMode struct {
	Mode
}

func (a *AbusoluteXMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_low := c.ram.Read(c.programCounter.Read())
	a.firstByte = &absolute_low
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_high := c.ram.Read(c.programCounter.Read())
	a.secondByte = &absolute_high
	index := c.x.Read()
	// why this calculation does not take 1 clock?
	absolute_low = absolute_low + index
	carry_out := absolute_low > 255
	if carry_out {
		// tick one more clock but no need to Increment
		// absolute_high because carry is included in low
		c.clock.Tick()
	} else {
		if c.IsWaitOneClock() {
			c.clock.Tick()
		}

	}
	address := (uint16(absolute_high<<8) + uint16(absolute_low)) & 0xffff
	return &address

}

func (a *AbusoluteXMode) GetModeString() string {
	return "AbusoluteX"
}

type AbusoluteYMode struct {
	Mode
}

func (a *AbusoluteYMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_low := c.ram.Read(c.programCounter.Read())
	a.firstByte = &absolute_low
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_high := c.ram.Read(c.programCounter.Read())
	a.secondByte = &absolute_high
	index := c.y.Read()
	// why this calculation does not take 1 clock?
	absolute_low = absolute_low + index
	carry_out := absolute_low > 255
	if carry_out {
		// tick one more clock but no need to Increment
		// absolute_high because carry is included in low
		c.clock.Tick()
	} else {
		if c.IsWaitOneClock() {
			c.clock.Tick()
		}
	}

	address := (uint16(absolute_high<<8) + uint16(absolute_low)) & 0xffff
	return &address

}
func (a *AbusoluteYMode) GetModeString() string {
	return "AbusoluteY"
}

type ZeroPageMode struct {
	Mode
}

func (z *ZeroPageMode) GetAddress(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	address_low := c.ram.Read(c.programCounter.Read())
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
	c.programCounter.Increment()
	zero_page_address := c.ram.Read(c.programCounter.Read())
	z.firstByte = &zero_page_address
	c.clock.Tick()
	// ignore carry
	zero_page_x_address := uint16((zero_page_address + c.x.Read()) & 255)
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
	c.programCounter.Increment()
	zero_page_address := c.ram.Read(c.programCounter.Read())
	z.firstByte = &zero_page_address
	c.clock.Tick()
	zero_page_y_address := uint16((zero_page_address + c.y.Read()) & 255)
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
	c.programCounter.Increment()
	relative := c.ram.Read(c.programCounter.Read())
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
	c.programCounter.Increment()
	indirect_address_low := c.ram.Read(c.programCounter.Read())
	i.firstByte = &indirect_address_low
	c.clock.Tick()
	c.programCounter.Increment()
	indirect_address_high := c.ram.Read(c.programCounter.Read())
	i.secondByte = &indirect_address_high
	c.clock.Tick()
	address_low := c.ram.Read(uint16(indirect_address_high<<8) + uint16(indirect_address_low))
	c.clock.Tick()
	// ignore carry from lower bit inrement
	address_high := c.ram.Read(uint16(indirect_address_high<<8) + uint16(indirect_address_low+1))
	address := (uint16(address_high<<8) + uint16(address_low)) & 0xffff
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
	c.programCounter.Increment()
	zero_page_address := c.ram.Read(c.programCounter.Read())
	i.firstByte = &zero_page_address
	// ignore carry_out so this shall lower than 0xff
	c.clock.Tick()
	indirect_address := (zero_page_address + c.x.Read()) & 0xff
	c.clock.Tick()
	address_low := c.ram.Read(uint16(indirect_address))
	c.clock.Tick()
	address_high := c.ram.Read(uint16((indirect_address + 1) & 0xff))
	address := (uint16(address_high<<8) + uint16(address_low)) & 0xffff
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
	c.programCounter.Increment()
	indirect_address := c.ram.Read(c.programCounter.Read())
	i.firstByte = &indirect_address
	c.clock.Tick()
	address_low := c.ram.Read(uint16(indirect_address))
	c.clock.Tick()
	address_high := c.ram.Read(uint16((indirect_address + 1) & 0xff))
	index := c.y.Read()
	address_low = address_low + index
	carry_out := address_low > 255
	if carry_out {
		c.clock.Tick()
	} else {
		if c.IsWaitOneClock() {
			c.clock.Tick()
		}
	}
	address := (uint16(address_high<<8) + uint16(address_low)) & 0xffff
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
