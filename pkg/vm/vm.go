package vm

import (
	"fmt"
)

type register struct {
	val uint8
}

type VM struct {
	Memory  []uint8
	memsize uint
	pc      uint
	rax     register
}

type instruction uint8

const (
	NOP instruction = iota
	MOV
	JMP
	CMP
	HLT
)

func New(memsize uint) (VM, error) {
	if memsize >= 1 {
		return VM{
			Memory:  make([]uint8, memsize),
			memsize: memsize,
			pc:      0,
			rax:     register{val: 0x45},
		}, nil
	} else {
		return VM{}, fmt.Errorf("give more memory lol")
	}
}

func (v *VM) Step() bool {
	if v.pc >= v.memsize {
		return false
	}

	var ci instruction = instruction(v.Memory[v.pc])
	switch ci {
	case NOP:
		fmt.Println("NOP")
	case MOV:
		fmt.Println("MOV")
	case JMP:
		fmt.Println("JMP")
	case CMP:
		fmt.Println("CMP")
	case HLT:
		fmt.Println("HLT")
		return false
	default:
		fmt.Printf("Unknown instruction 0x%x at 0x%x, halting\n", ci, v.pc)
		return false
	}

	v.pc += 1
	return true
}
