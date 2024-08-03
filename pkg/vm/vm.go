package vm

import (
	"fmt"
)

// 64 kib
const MEMSIZE = 65536

type VM struct {
	Memory   [MEMSIZE]uint8
	regs     [NREGS]register
	print_bs bool
}

func New(bs bool) VM {
	regs := [NREGS]register{}

	for i := 0; i < NREGS; i++ {
		regs[i].kind = reg(i)
		regs[i].val = 0
	}

	regs[BP].val = MEMSIZE - 2
	regs[SP].val = MEMSIZE - 2

	return VM{
		Memory:   [MEMSIZE]uint8{},
		regs:     regs,
		print_bs: bs,
	}
}

func (v *VM) SetReg(r reg, val uint16) {
	v.regs[r].val = val
}

func (v *VM) GetReg(r reg) uint16 {
	return v.regs[r].val
}

func (v *VM) inc_pc(a uint16) {
	v.regs[PC].val += a
}

func (v *VM) set_pc(a uint16) {
	v.regs[PC].val = a
}

func (v *VM) get_pc() uint16 {
	return v.regs[PC].val
}

func (v *VM) Step() bool {
	if v.get_pc() >= MEMSIZE-1 {
		return false
	}

	var ci opcode = opcode(v.Memory[v.get_pc()])

	handle, ok := handler[ci]
	if ok {
		return handle.do(v)
	} else {
		fmt.Printf("unregistered opcode %d\n", ci)
		return false
	}
}
