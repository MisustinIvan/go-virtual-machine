package vm

import "fmt"

const STACK_SIZE = 1024

type pus_op struct{}

func (o pus_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	if v.regs[BP].val-v.regs[SP].val >= STACK_SIZE {
		if v.print_bs {
			fmt.Println("stack overflow")
		}
		return false
	}

	r0 := v.regs[v.Memory[v.get_pc()+1]]

	if v.print_bs {
		fmt.Printf("0x%x:PUS 0x%x %s:0x%x\n", v.get_pc(), v.GetReg(SP), regtostring(r0.kind), r0.val)
	}

	ok := v.writeu16(v.GetReg(SP), r0.val)
	if !ok {
		fmt.Println("wtf2")
		return false
	}

	v.SetReg(SP, v.GetReg(SP)-2)

	v.inc_pc(uint16(o.size()))
	return true
}

func (o pus_op) size() uint8 {
	return 2
}

type pop_op struct{}

func (o pop_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	if v.GetReg(SP)+2 > v.GetReg(BP) {
		return false
	}

	r0 := v.regs[v.Memory[v.get_pc()+1]]

	val, ok := v.readu16(v.GetReg(SP) + 2)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:POP %s 0x%x:0x%x\n", v.get_pc(), regtostring(r0.kind), v.GetReg(SP), val)
	}

	v.SetReg(r0.kind, val)
	v.SetReg(SP, v.GetReg(SP)+2)

	v.inc_pc(uint16(o.size()))
	return true
}

func (o pop_op) size() uint8 {
	return 2
}
