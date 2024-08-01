package vm

import (
	"fmt"
)

// one less than 65536
const MEMSIZE = 65535
const NREGS = 8

type register struct {
	val  uint16
	kind reg
}

type VM struct {
	Memory [MEMSIZE]uint8
	pc     uint16
	regs   [NREGS]register
}

type reg uint8

const (
	PC reg = iota
	BP
	SP
	RA
	RB
	RC
	RD
	RS
)

type instruction uint8

const (
	NOP instruction = iota
	ADD
	SUB
	MUL
	DIV
	MOV
	JMP
	INC
	DEC
	LOD
	STR
	PRT
	HLT
	CMP
	JIO
	JIZ
	JIE
	JNE
	JIG
	JIS
	JEG
	JES
	XOR
)

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
	if v.pc >= MEMSIZE {
		return false
	}

	var ci instruction = instruction(v.Memory[v.pc])
	switch ci {
	case NOP:
		fmt.Printf("0x%x:NOP\n", v.pc)
		v.pc += 1

	case ADD:
		if v.pc+3 >= MEMSIZE {
			return false
		} else {
			v.add(&v.regs[v.Memory[v.pc+1]], &v.regs[v.Memory[v.pc+2]])
		}

	case SUB:
		if v.pc+3 >= MEMSIZE {
			return false
		} else {
			v.sub(&v.regs[v.Memory[v.pc+1]], &v.regs[v.Memory[v.pc+2]])
		}

	case MUL:
		if v.pc+3 >= MEMSIZE {
			return false
		} else {
			v.mul(&v.regs[v.Memory[v.pc+1]], &v.regs[v.Memory[v.pc+2]])
		}

	case DIV:
		if v.pc+3 >= MEMSIZE {
			return false
		} else {
			v.div(&v.regs[v.Memory[v.pc+1]], &v.regs[v.Memory[v.pc+2]])
		}

	case MOV:
		if v.pc+3 >= MEMSIZE {
			return false
		} else {
			v.mov(&v.regs[v.Memory[v.pc+1]], &v.regs[v.Memory[v.pc+2]])
		}

	case LOD:
		if v.pc+4 >= MEMSIZE {
			return false
		} else {
			v.lod(&v.regs[v.Memory[v.pc+1]], v.readu16(v.pc+2))
		}

	case STR:
		if v.pc+4 >= MEMSIZE {
			return false
		} else {
			v.str(v.readu16(v.pc+2), v.regs[v.Memory[v.pc+1]])
		}

	case PRT:
		if v.pc+2 >= MEMSIZE {
			return false
		} else {
			v.prt(v.regs[v.Memory[v.pc+1]])
		}

	case HLT:
		fmt.Printf("0x%x:HLT\n", v.pc)
		return false

	default:
		fmt.Printf("Unknown instruction 0x%x at 0x%x, halting\n", ci, v.pc)
		return false
	}
	return true
}
