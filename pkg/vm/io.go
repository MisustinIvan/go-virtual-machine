package vm

import "fmt"

type prt_op struct{}

func (o prt_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.pc+1]]

	if v.print_bs {
		fmt.Printf("0x%x:PRT %s:0x%x:%d\n", v.pc, regtostring(r0.kind), r0.val, r0.val)
	}

	fmt.Println(r0.val)

	v.inc_pc(uint16(o.size()))
	return true
}

func (o prt_op) size() uint8 {
	return 2
}

type pr8_op struct{}

func (o pr8_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.pc+1]]

	if v.print_bs {
		fmt.Printf("0x%x:PRT %s:0x%x:%d\n", v.pc, regtostring(r0.kind), uint8(r0.val), uint8(r0.val))
	}

	fmt.Printf("%c", uint8(r0.val))

	v.inc_pc(uint16(o.size()))
	return true
}

func (o pr8_op) size() uint8 {
	return 2
}
