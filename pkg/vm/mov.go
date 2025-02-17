package vm

import "fmt"

type mov_op struct{}

func (o mov_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]
	var r1 *register = &v.regs[v.Memory[v.get_pc()+2]]

	if v.print_bs {
		fmt.Printf("0x%x:MOV %s:0x%x %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	}
	r0.val = r1.val

	v.inc_pc(uint16(o.size()))
	return true
}

func (o mov_op) size() uint8 {
	return 3
}

type mvi_op struct{}

func (o mvi_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]
	val, ok := v.readu16(v.get_pc() + 2)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:MVI %s:0x%x 0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, val)
	}
	r0.val = val

	v.inc_pc(uint16(o.size()))
	return true
}

func (o mvi_op) size() uint8 {
	return 4
}

type lod_op struct{}

func (o lod_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]

	address, ok := v.readu16(v.get_pc() + 2)
	if !ok {
		return false
	}

	val, ok := v.readu16(address)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:LOD %s:0x%x 0x%x:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, address, val)
	}
	r0.val = val

	v.inc_pc(uint16(o.size()))
	return true
}

func (o lod_op) size() uint8 {
	return 4
}

type ld8_op struct{}

func (o ld8_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]

	address, ok := v.readu16(v.get_pc() + 2)
	if !ok {
		return false
	}

	val, ok := v.readu8(address)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:LD8 %s:0x%x 0x%x:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, address, val)
	}
	r0.val = val

	v.inc_pc(uint16(o.size()))
	return true
}

func (o ld8_op) size() uint8 {
	return 4
}

type str_op struct{}

func (o str_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	var r0 = v.regs[v.Memory[v.get_pc()+3]]

	if v.print_bs {
		fmt.Printf("0x%x:STR 0x%x:0x%x %s:0x%x\n", v.get_pc(), address, r0.val, regtostring(r0.kind), r0.val)
	}

	ok = v.writeu16(address, r0.val)
	if !ok {
		return false
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o str_op) size() uint8 {
	return 4
}

type st8_op struct{}

func (o st8_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	val, ok := v.readu16(address)
	if !ok {
		return false
	}

	var r0 = v.regs[v.Memory[v.get_pc()+3]]

	if v.print_bs {
		fmt.Printf("0x%x:ST8 0x%x:0x%x %s:0x%x\n", v.get_pc(), address, val, regtostring(r0.kind), r0.val)
	}

	ok = v.writeu8(address, r0.val)
	if !ok {
		return false
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o st8_op) size() uint8 {
	return 4
}
