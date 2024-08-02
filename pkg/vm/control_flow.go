package vm

import "fmt"

type cmp_op struct{}

func (o cmp_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var a = v.regs[v.Memory[v.pc+1]].val
	var b = v.regs[v.Memory[v.pc+1]].val

	if a == b {
		v.regs[SP].val = uint16(EQUALS)
	} else if a < b {
		v.regs[SP].val = uint16(SMALLER)
	} else {
		v.regs[SP].val = uint16(GREATER)
	}

	v.pc += uint16(o.size())
	return true
}

func (o cmp_op) size() uint8 {
	return 3
}

type jif_op struct{}

func (o jif_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(OVERFLOW) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jif_op) size() uint8 {
	return 3
}

type jid_op struct{}

func (o jid_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(DIVZERO) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jid_op) size() uint8 {
	return 3
}

type jiz_op struct{}

func (o jiz_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(ZERO) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jiz_op) size() uint8 {
	return 3
}

type jie_op struct{}

func (o jie_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(EQUALS) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jie_op) size() uint8 {
	return 3
}

type jne_op struct{}

func (o jne_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val == uint16(EQUALS) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jne_op) size() uint8 {
	return 3
}

type jig_op struct{}

func (o jig_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(GREATER) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jig_op) size() uint8 {
	return 3
}

type jis_op struct{}

func (o jis_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(SMALLER) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jis_op) size() uint8 {
	return 3
}

type jeg_op struct{}

func (o jeg_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(GREATER) && v.regs[SP].val != uint16(EQUALS) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jeg_op) size() uint8 {
	return 3
}

type jes_op struct{}

func (o jes_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	if v.regs[SP].val != uint16(SMALLER) && v.regs[SP].val != uint16(EQUALS) {
		v.pc += uint16(o.size())
		return true
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jes_op) size() uint8 {
	return 3
}

type jmp_op struct{}

func (o jmp_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
		return false
	}

	// cant jump to pc
	if address == v.pc {
		return false
	}

	fmt.Printf("0x%x:JMP 0x%d\n", v.pc, address)

	v.pc = address

	return true
}

func (o jmp_op) size() uint8 {
	return 3
}

type nop_op struct{}

func (o nop_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	fmt.Printf("0x%x:NOP\n", v.pc)

	v.pc += uint16(o.size())
	return true
}

func (o nop_op) size() uint8 {
	return 1
}

type hlt_op struct{}

func (o hlt_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	fmt.Printf("0x%x:HLT\n", v.pc)

	v.pc += uint16(o.size())
	return false
}

func (o hlt_op) size() uint8 {
	return 1
}
