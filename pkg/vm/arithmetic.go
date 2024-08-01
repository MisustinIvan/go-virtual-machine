package vm

import "fmt"

func (v *VM) add(r0 *register, r1 *register) {
	fmt.Printf("0x%x:ADD %s:0x%x %s:0x%x\n", v.pc, regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	v.regs[v.Memory[v.pc+1]].val = r0.val + r1.val
	if int(r0.val)+int(r1.val) >= MEMSIZE {
		v.regs[SP].val = uint16(OVERFLOW)
	}
	v.pc += 3
}

func (v *VM) sub(r0 *register, r1 *register) {
	fmt.Printf("0x%x:SUB %s:0x%x %s:0x%x\n", v.pc, regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	v.regs[v.Memory[v.pc+1]].val = r0.val - r1.val
	if int(r0.val)-int(r1.val) < 0 {
		v.regs[SP].val = uint16(OVERFLOW)
	}
	v.pc += 3
}

func (v *VM) mul(r0 *register, r1 *register) {
	fmt.Printf("0x%x:MUL %s:0x%x %s:0x%x\n", v.pc, regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	fmt.Println(r0.val * r1.val)
	v.regs[v.Memory[v.pc+1]].val = r0.val * r1.val
	if int(r0.val)*int(r1.val) >= MEMSIZE {
		v.regs[SP].val = uint16(OVERFLOW)
	}
	v.pc += 3
}

func (v *VM) div(r0 *register, r1 *register) {
	fmt.Printf("0x%x:DIV %s:0x%x %s:0x%x\n", v.pc, regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	if r1.val == 0 {
		fmt.Println("division by zero")
		v.regs[SP].val = uint16(DIVZERO)
	} else {
		fmt.Println(r0.val / r1.val)
		v.regs[v.Memory[v.pc+1]].val = r0.val / r1.val
	}
	v.pc += 3
}
