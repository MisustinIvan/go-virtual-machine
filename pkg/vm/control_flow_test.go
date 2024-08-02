package vm_test

import (
	"main/pkg/vm"
	"testing"
)

func TestJmp(t *testing.T) {
	machine := vm.New(false)

	var val uint16 = 6942

	machine.Memory[0] = uint8(vm.JMP)
	machine.Memory[1] = uint8(val)
	machine.Memory[2] = uint8(val >> 8)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	t.Log(machine.GetReg(vm.PC))

	if machine.GetReg(vm.PC) != val {
		t.Fatal("did not jump to correct address")
	}
}

func TestEndJmp(t *testing.T) {
	machine := vm.New(false)

	machine.Memory[0] = uint8(vm.JMP)
	machine.Memory[1] = 0xff
	machine.Memory[2] = 0xff

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	t.Log(machine.GetReg(vm.PC))

	if machine.Step() {
		t.Fatal("machine did not stop as supposed")
	}
	t.Log("next step caused halt as it should")
}

func TestPcJmp(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.PC, 69)

	machine.Memory[69] = uint8(vm.JMP)
	machine.Memory[70] = 0x45 // hehe
	machine.Memory[71] = 0

	if machine.Step() {
		t.Fatal("did not stop machine!")
	}
}
