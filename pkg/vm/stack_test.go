package vm_test

import (
	"go-virtual-machine/pkg/vm"
	"testing"
)

func TestPus(t *testing.T) {
	machine := vm.New(false)
	number := uint16(69)

	machine.SetReg(vm.RA, number)

	machine.Memory[0] = uint8(vm.PUS)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.PUS)
	machine.Memory[3] = uint8(vm.RA)

	t.Logf("SP: %d\n", machine.GetReg(vm.SP))
	t.Logf("stepping\n")

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	stored_l := machine.Memory[machine.GetReg(vm.BP)]
	stored_h := machine.Memory[machine.GetReg(vm.BP)+1]

	stored := uint16(stored_l) | (uint16(stored_h) << 8)

	t.Logf("Stored lower: 0x%x\n", stored_l)
	t.Logf("Stored higher: 0x%x\n", stored_h)
	t.Logf("Stored value: 0x%x\n", stored)

	t.Logf("SP: %d", machine.GetReg(vm.SP))
	t.Logf("PC: 0x%x", machine.GetReg(vm.PC))

	if stored != number {
		t.Fatal("wrong stored number")
	}

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	stored_l = machine.Memory[machine.GetReg(vm.BP)-2]
	stored_h = machine.Memory[machine.GetReg(vm.BP)-1]

	stored = uint16(stored_l) | (uint16(stored_h) << 8)

	t.Logf("Stored lower: 0x%x\n", stored_l)
	t.Logf("Stored higher: 0x%x\n", stored_h)
	t.Logf("Stored value: 0x%x\n", stored)

	t.Logf("SP: %d", machine.GetReg(vm.SP))
	t.Logf("PC: 0x%x", machine.GetReg(vm.PC))

	if stored != number {
		t.Fatal("wrong stored number")
	}
}

func TestPop(t *testing.T) {
	machine := vm.New(false)
	number := uint16(69)

	machine.SetReg(vm.RA, number)

	machine.Memory[0] = uint8(vm.PUS)
	machine.Memory[1] = uint8(vm.RA)

	machine.Memory[2] = uint8(vm.POP)
	machine.Memory[3] = uint8(vm.RB)

	machine.Memory[4] = uint8(vm.CMP)
	machine.Memory[5] = uint8(vm.RA)
	machine.Memory[6] = uint8(vm.RB)

	// not in a loop to know which instruction caused it
	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}
	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}
	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.RA) != machine.GetReg(vm.RB) && machine.GetReg(vm.RS) != uint16(vm.EQUALS) {
		t.Fatal("failed to pop")
	}
}
