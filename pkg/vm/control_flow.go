package vm

import "fmt"

type jmp_op struct{}

func (o jmp_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.pc + 1)
	if !ok {
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
