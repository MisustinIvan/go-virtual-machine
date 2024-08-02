package vm_test

import (
	"main/pkg/vm"
	"testing"
)

func TestAdd(t *testing.T) {
	machine := vm.New(false)

	var n0 uint16 = 69
	var n1 uint16 = 420

	machine.SetReg(vm.RA, n0)
	machine.SetReg(vm.RB, n1)

	machine.Memory[0] = uint8(vm.ADD)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.RA) != n0+n1 {
		t.Fatal("instruction returned wrong result")
	}
}
