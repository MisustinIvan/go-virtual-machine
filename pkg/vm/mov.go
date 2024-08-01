package vm

import "fmt"

func (v *VM) mov(r0 *register, r1 *register) {
	fmt.Printf("0x%x:MOV %s:0x%x %s:0x%x\n", v.pc, regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	r0.val = r1.val
	v.pc += 3
}

func (v *VM) lod(r0 *register, address uint16) {
	var lower uint8 = v.Memory[address]
	var higher uint8 = v.Memory[address+1]
	var val uint16 = uint16(lower) | (uint16(higher) << 8)
	fmt.Printf("0x%x:LOD %s:0x%x 0x%x:0x%x\n", v.pc, regtostring(r0.kind), r0.val, address, val)
	r0.val = val

	v.pc += 4
}

func (v *VM) str(address uint16, r0 register) {
	var lower uint8 = v.Memory[address]
	var higher uint8 = v.Memory[address+1]
	var val uint16 = uint16(lower) | (uint16(higher) << 8)
	fmt.Printf("0x%x:STR 0x%x:0x%x %s:0x%x\n", v.pc, address, val, regtostring(r0.kind), r0.val)
	r0.val = val

	v.pc += 4
}

func (v *VM) readu16(addres uint16) uint16 {
	var lower uint8 = v.Memory[addres]
	var higher uint8 = v.Memory[addres+1]
	return uint16(lower) | (uint16(higher) << 8)
}
