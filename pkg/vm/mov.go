package vm

import "fmt"

type mov_op struct{}

func (o mov_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.pc+1]]
	var r1 *register = &v.regs[v.Memory[v.pc+2]]

	fmt.Printf("0x%x:MOV %s:0x%x %s:0x%x\n", v.pc, regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	r0.val = r1.val

	v.pc += uint16(o.size())
	return true
}

func (o mov_op) size() uint8 {
	return 3
}

type lod_op struct{}

func (o lod_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.pc+1]]

	address, ok := v.readu16(v.pc + 2)
	if !ok {
		return false
	}

	val, ok := v.readu16(address)
	if !ok {
		return false
	}

	fmt.Printf("0x%x:LOD %s:0x%x 0x%x:0x%x\n", v.pc, regtostring(r0.kind), r0.val, address, val)
	r0.val = val

	v.pc += uint16(o.size())
	return true
}

func (o lod_op) size() uint8 {
	return 4
}

type str_op struct{}

func (o str_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(uint16(v.Memory[v.pc+1]))
	if !ok {
		return false
	}

	var r0 = v.regs[v.Memory[v.pc+3]]

	// lets comment this out...
	// var lower uint8 = uint8(r0.val >> 8)
	// var higher uint8 = uint8(r0.val)
	// var val uint16 = uint16(lower) | (uint16(higher) << 8)
	fmt.Printf("0x%x:STR 0x%x:0x%x %s:0x%x\n", v.pc, address, r0.val, regtostring(r0.kind), r0.val)

	ok = v.writeu16(address, r0.val)
	if !ok {
		return false
	}

	v.pc += uint16(o.size())
	return true
}

func (o str_op) size() uint8 {
	return 4
}
