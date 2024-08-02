package vm_test

import (
	"main/pkg/vm"
	"testing"
)

func TestMov(t *testing.T) {
	machine := vm.New()

	machine.SetReg(vm.RA, 69)

	machine.Memory[0] = uint8(vm.MOV)
	machine.Memory[1] = uint8(vm.RB)
	machine.Memory[2] = uint8(vm.RA)

	if !machine.Step() {
		t.Fatal("could not execute mov instruction")
	}

	if machine.GetReg(vm.RA) != machine.GetReg(vm.RB) || machine.GetReg(vm.RA) != 69 {
		t.Fatal("mov instruction failed")
	}
}
