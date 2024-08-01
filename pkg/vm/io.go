package vm

import "fmt"

func (v *VM) prt(r0 register) {
	fmt.Printf("0x%x:PRT %s:0x%x:%d\n", v.pc, regtostring(r0.kind), r0.val, r0.val)
	v.pc += 2
}
