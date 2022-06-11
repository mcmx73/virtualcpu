package virtualcpu

import (
	"fmt"
	"sync"
)

const (
	REG_A = iota
	REG_B
	REG_C
	REG_D
	REG_H
	REG_L

	FLAG_ZF = 0xF + iota
	FLAG_CF
)

type CpuCore struct {
	mux sync.Mutex
	A   ByteRegister
	B   ByteRegister
	C   ByteRegister
	D   ByteRegister
	H   ByteRegister
	L   ByteRegister
	//
	// Program Counter, point to current memory position
	PC WordRegister
	// Stack Pointer, show current stack position
	SP WordRegister
	// Flags
	ZF Flag
	CF Flag
}

type RegID int8
type FlagID int8

type Option func(c *CpuCore)

func (c *CpuCore) Init(options ...Option) {
	for _, opt := range options {
		opt(c)
	}
}

func (c *CpuCore) Dump() string {
	return "IMPLEMENT ME"
}

func (c *CpuCore) setRegister(register RegID, data byte) {
	switch register {
	case REG_A:
		c.A.Set(data)
	case REG_B:
		c.B.Set(data)
	case REG_C:
		c.C.Set(data)
	case REG_D:
		c.D.Set(data)
	case REG_H:
		c.H.Set(data)
	case REG_L:
		c.L.Set(data)
	default:
		panic(fmt.Sprintf("Unknown Register ID: %v", register))
	}
}

func (c *CpuCore) getRegister(register RegID) (data byte) {
	switch register {
	case REG_A:
		return c.A.Get()
	case REG_B:
		return c.B.Get()
	case REG_C:
		return c.C.Get()
	case REG_D:
		return c.D.Get()
	case REG_H:
		return c.H.Get()
	case REG_L:
		return c.L.Get()
	default:
		panic(fmt.Sprintf("Unknown Register ID: %v", register))
	}
	return 0
}

func (c *CpuCore) setFlag(flag FlagID, value bool) {
	switch flag {
	case FLAG_ZF:
		c.ZF.Set(value)
	case FLAG_CF:
		c.CF.Set(value)
	default:
		panic(fmt.Sprintf("Unknown Flag ID: %v", flag))
	}
}

func (c *CpuCore) getFlag(flag FlagID) (value bool) {
	switch flag {
	case FLAG_ZF:
		return c.ZF.Get()
	case FLAG_CF:
		return c.CF.Get()
	default:
		panic(fmt.Sprintf("Unknown Flag ID: %v", flag))
	}
	return false
}
