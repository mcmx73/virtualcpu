package virtualcpu

func SetByteRegistersDefault() Option {
	return func(c *CpuCore) {
		c.A = &byteRegister{
			value: 0,
		}
		c.B = &byteRegister{
			value: 0,
		}
		c.C = &byteRegister{
			value: 0,
		}
		c.D = &byteRegister{
			value: 0,
		}
		c.H = &byteRegister{
			value: 0,
		}
		c.L = &byteRegister{
			value: 0,
		}
	}
}

func SetMemoryDefault() Option {
	return func(c *CpuCore) {
		c.memory = &MemoryDefault{}
	}
}
