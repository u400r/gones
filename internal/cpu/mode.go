package cpu

func ImmediateMode(c *Cpu) *uint16 {
	if false {
		c.clock.Tick()
	}
	c.programCounter.Increment()
	address := c.programCounter.Read()
	return &address

}

func AbusoluteMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_low := c.ram.Read(c.programCounter.Read())
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_high := c.ram.Read(c.programCounter.Read())
	address := uint16(absolute_high<<8) + uint16(absolute_low)
	return &address

}

func AbusoluteXMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_low := c.ram.Read(c.programCounter.Read())
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_high := c.ram.Read(c.programCounter.Read())
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

func AbusoluteYMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_low := c.ram.Read(c.programCounter.Read())
	c.clock.Tick()
	c.programCounter.Increment()
	absolute_high := c.ram.Read(c.programCounter.Read())
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

func ZeroPageMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	address_low := uint16(c.ram.Read(c.programCounter.Read()))
	return &address_low

}

func ZeroPageXMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	zero_page_address := c.ram.Read(c.programCounter.Read())
	c.clock.Tick()
	// ignore carry
	zero_page_x_address := uint16((zero_page_address + c.x.Read()) & 255)
	return &zero_page_x_address
}

func ZeroPageYMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	zero_page_address := c.ram.Read(c.programCounter.Read())
	c.clock.Tick()
	zero_page_y_address := uint16((zero_page_address + c.y.Read()) & 255)
	return &zero_page_y_address

}

func RelativeMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	relative := uint16(c.ram.Read(c.programCounter.Read()))
	return &relative

}

func IndirectMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	indirect_address_low := c.ram.Read(c.programCounter.Read())
	c.clock.Tick()
	c.programCounter.Increment()
	indirect_address_high := c.ram.Read(c.programCounter.Read())
	c.clock.Tick()
	address_low := c.ram.Read(uint16(indirect_address_high<<8) + uint16(indirect_address_low))
	c.clock.Tick()
	// ignore carry from lower bit inrement
	address_high := c.ram.Read(uint16(indirect_address_high<<8) + uint16(indirect_address_low+1))
	address := (uint16(address_high<<8) + uint16(address_low)) & 0xffff
	return &address

}

func IndirectXMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	zero_page_address := c.ram.Read(c.programCounter.Read())
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

func IndirectYMode(c *Cpu) *uint16 {
	c.clock.Tick()
	c.programCounter.Increment()
	indirect_address := c.ram.Read(c.programCounter.Read())
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

func ImplicitMode(c *Cpu) *uint16 {
	if false {
		c.clock.Tick()
	}
	return nil
}
