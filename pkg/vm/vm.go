package vm

import (
	"fmt"
)

// 64 kib
const MEMSIZE = 65536

type VM struct {
	Memory [MEMSIZE]uint8
	pc     uint16
	regs   [NREGS]register
}

type sp_code uint16

const (
	ZERO sp_code = iota
	ONE
	OVERFLOW
	DIVZERO
	EQUALS
	NOT_EQUALS
	GREATER
	SMALLER
	GREATER_OR_EQUAL
	SMALLER_OR_EQUAL
)

func New() VM {
	regs := [NREGS]register{}

	for i := 0; i < NREGS; i++ {
		regs[i].kind = reg(i)
		regs[i].val = 0
	}

	return VM{
		Memory: [MEMSIZE]uint8{},
		pc:     0,
		regs:   regs,
	}
}

func (v *VM) Step() bool {
	if v.pc >= MEMSIZE-1 {
		return false
	}

	var ci opcode = opcode(v.Memory[v.pc])

	handle, ok := handler[ci]
	if ok {
		return handle.do(v)
	} else {
		fmt.Printf("unregistered opcode %d\n", ci)
		return false
	}
}
