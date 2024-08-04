package vm

import "fmt"

type cmp_op struct{}

func (o cmp_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var r0 = v.regs[v.Memory[v.get_pc()+1]]
	var r1 = v.regs[v.Memory[v.get_pc()+2]]

	if r0.val == r1.val {
		v.regs[RS].val = uint16(EQUALS)
	} else if r0.val < r1.val {
		v.regs[RS].val = uint16(SMALLER)
	} else {
		v.regs[RS].val = uint16(GREATER)
	}

	if v.print_bs {
		fmt.Printf("0x%x:CMP %s:0x%x %s:0x%x\n", v.get_pc(), regtostring(r0.kind), r0.val, regtostring(r1.kind), r1.val)
	}

	v.inc_pc(uint16(o.size()))
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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(OVERFLOW) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JIF 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(DIVZERO) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JID 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(ZERO) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JIZ 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(EQUALS) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JIE 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val == uint16(EQUALS) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JNE 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(GREATER) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JIG 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(SMALLER) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JIS 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(GREATER) && v.regs[RS].val != uint16(EQUALS) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JEG 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	if v.regs[RS].val != uint16(SMALLER) && v.regs[RS].val != uint16(EQUALS) {
		v.inc_pc(uint16(o.size()))
		return true
	}

	if v.print_bs {
		fmt.Printf("0x%x:JES 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

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

	address, ok := v.readu16(v.get_pc() + 1)
	if !ok {
		return false
	}

	// cant jump to pc
	if address == v.get_pc() {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:JMP 0x%d\n", v.get_pc(), address)
	}

	v.set_pc(address)

	return true
}

func (o jmp_op) size() uint8 {
	return 3
}

type jpi_op struct{}

func (o jpi_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	var address uint16
	if v.Memory[v.get_pc()+1] >= NREGS {
		return false
	} else {
		address = v.regs[v.Memory[v.get_pc()+1]].val
	}
	// cant jump to pc
	if address == v.get_pc() {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:JPI %s:0x%d\n", v.get_pc(), regtostring(v.regs[v.Memory[v.get_pc()+1]].kind), address)
	}

	v.set_pc(address)

	return true
}

func (o jpi_op) size() uint8 {
	return 2
}

type nop_op struct{}

func (o nop_op) do(v *VM) bool {
	if !op_ok(o, *v) {
		return false
	}

	if v.print_bs {
		fmt.Printf("0x%x:NOP\n", v.get_pc())
	}

	v.inc_pc(uint16(o.size()))
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

	if v.print_bs {
		fmt.Printf("0x%x:HLT\n", v.get_pc())
	}

	v.inc_pc(uint16(o.size()))
	return false
}

func (o hlt_op) size() uint8 {
	return 1
}
