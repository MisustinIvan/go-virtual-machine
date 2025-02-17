package vm

import "fmt"

type add_op struct{}

func (o add_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.get_pc()+1]]
	var r1 register = v.regs[v.Memory[v.get_pc()+2]]

	if v.print_bs {
		fmt.Printf("0x%x:ADD %s:0x%x %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	}

	v.regs[v.Memory[v.get_pc()+1]].val = r0.val + r1.val

	if int(r0.val)+int(r1.val) >= MEMSIZE {
		v.regs[RS].val = uint16(OVERFLOW)
	} else if r0.val == 0 {
		v.regs[RS].val = uint16(ZERO)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o add_op) size() uint8 {
	return 3
}

type adi_op struct{}

func (o adi_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.get_pc()+1]]
	val, ok := v.readu16(v.get_pc() + 2)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:ADI %s:0x%x 0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, val)
	}

	v.regs[v.Memory[v.get_pc()+1]].val = r0.val + val

	if int(r0.val)+int(val) >= MEMSIZE {
		v.regs[RS].val = uint16(OVERFLOW)
	} else if r0.val == 0 {
		v.regs[RS].val = uint16(ZERO)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o adi_op) size() uint8 {
	return 4
}

type sub_op struct{}

func (o sub_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.get_pc()+1]]
	var r1 register = v.regs[v.Memory[v.get_pc()+2]]

	if v.print_bs {
		fmt.Printf("0x%x:SUB %s:0x%x %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	}
	v.regs[v.Memory[v.get_pc()+1]].val = r0.val - r1.val
	if int(r0.val)-int(r1.val) < 0 {
		v.regs[RS].val = uint16(OVERFLOW)
	} else if r0.val == r1.val {
		v.regs[RS].val = uint16(ZERO)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o sub_op) size() uint8 {
	return 3
}

type sbi_op struct{}

func (o sbi_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.get_pc()+1]]
	val, ok := v.readu16(v.get_pc() + 2)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:SBI %s:0x%x 0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, val)
	}
	v.regs[v.Memory[v.get_pc()+1]].val = r0.val - val
	if int(r0.val)-int(val) < 0 {
		v.regs[RS].val = uint16(OVERFLOW)
	} else if r0.val == val {
		v.regs[RS].val = uint16(ZERO)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o sbi_op) size() uint8 {
	return 4
}

type mul_op struct{}

func (o mul_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]
	var r1 *register = &v.regs[v.Memory[v.get_pc()+2]]

	if v.print_bs {
		fmt.Printf("0x%x:MUL %s:0x%x %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	}

	v.regs[v.Memory[v.get_pc()+1]].val = r0.val * r1.val

	if int(r0.val)*int(r1.val) >= MEMSIZE {
		v.regs[RS].val = uint16(OVERFLOW)
	} else if r0.val == 0 || r1.val == 0 {
		v.regs[RS].val = uint16(ZERO)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o mul_op) size() uint8 {
	return 3
}

type mli_op struct{}

func (o mli_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]
	val, ok := v.readu16(v.get_pc() + 2)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:MLI %s:0x%x 0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, val)
	}

	v.regs[v.Memory[v.get_pc()+1]].val = r0.val * val

	if int(r0.val)*int(val) >= MEMSIZE {
		v.regs[RS].val = uint16(OVERFLOW)
	} else if r0.val == 0 || val == 0 {
		v.regs[RS].val = uint16(ZERO)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o mli_op) size() uint8 {
	return 4
}

type div_op struct{}

func (o div_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]
	var r1 *register = &v.regs[v.Memory[v.get_pc()+2]]

	if v.print_bs {
		fmt.Printf("0x%x:DIV %s:0x%x %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	}
	if r1.val == 0 {
		if v.print_bs {
			fmt.Println("division by zero")
		}
		v.regs[RS].val = uint16(DIVZERO)
		r0.val = 0
	} else {
		if v.print_bs {
			fmt.Println(r0.val / r1.val)
		}
		v.regs[v.Memory[v.get_pc()+1]].val = r0.val / r1.val

		if r0.val == 0 {
			v.regs[RS].val = uint16(ZERO)
		} else {
			v.regs[RS].val = uint16(NONE)
		}
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o div_op) size() uint8 {
	return 3
}

type dvi_op struct{}

func (o dvi_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 *register = &v.regs[v.Memory[v.get_pc()+1]]
	val, ok := v.readu16(v.get_pc() + 2)
	if !ok {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:DVI %s:0x%x 0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, val)
	}
	if val == 0 {
		if v.print_bs {
			fmt.Println("division by zero")
		}
		v.regs[RS].val = uint16(DIVZERO)
		r0.val = 0
	} else {
		if v.print_bs {
			fmt.Println(r0.val / val)
		}
		v.regs[v.Memory[v.get_pc()+1]].val = r0.val / val

		if r0.val == 0 {
			v.regs[RS].val = uint16(ZERO)
		} else {
			v.regs[RS].val = uint16(NONE)
		}
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o dvi_op) size() uint8 {
	return 4
}

type inc_op struct{}

func (o inc_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.get_pc()+1]]

	if v.print_bs {
		fmt.Printf("0x%x:INC %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val)
	}

	v.regs[v.Memory[v.get_pc()+1]].val += 1

	if int(r0.val)+1 >= MEMSIZE {
		v.regs[RS].val = uint16(OVERFLOW)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o inc_op) size() uint8 {
	return 2
}

type dec_op struct{}

func (o dec_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 register = v.regs[v.Memory[v.get_pc()+1]]

	if v.print_bs {
		fmt.Printf("0x%x:DEC %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val)
	}

	v.regs[v.Memory[v.get_pc()+1]].val -= 1

	if int(r0.val)-1 < 0 {
		v.regs[RS].val = uint16(OVERFLOW)
	} else if int(r0.val)-1 == 0 {
		v.regs[RS].val = uint16(ZERO)
	} else {
		v.regs[RS].val = uint16(NONE)
	}

	v.inc_pc(uint16(o.size()))
	return true
}

func (o dec_op) size() uint8 {
	return 2
}
