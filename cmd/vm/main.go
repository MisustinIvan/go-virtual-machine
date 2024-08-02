package main

import (
	"main/pkg/vm"
)

func main() {
	machine := vm.New()

	// LOD RA 69
	machine.Memory[0] = uint8(vm.LOD)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(69)
	machine.Memory[3] = uint8(0)

	// MOV RB RA
	machine.Memory[4] = uint8(vm.MOV)
	machine.Memory[5] = uint8(vm.RB)
	machine.Memory[6] = uint8(vm.RA)

	// INC RB
	machine.Memory[7] = uint8(vm.INC)
	machine.Memory[8] = uint8(vm.RB)

	// STR 71 RB
	machine.Memory[9] = uint8(vm.STR)
	machine.Memory[10] = uint8(71)
	machine.Memory[11] = uint8(0)
	machine.Memory[12] = uint8(vm.RB)

	// PRT RA
	machine.Memory[13] = uint8(vm.PRT)
	machine.Memory[14] = uint8(vm.RA)
	// PRT RB
	machine.Memory[15] = uint8(vm.PRT)
	machine.Memory[16] = uint8(vm.RB)

	machine.Memory[17] = uint8(vm.JMP)
	machine.Memory[18] = uint8(7)
	machine.Memory[19] = uint8(0)
	// HLT
	machine.Memory[20] = uint8(vm.HLT)

	// constant data
	machine.Memory[69] = uint8(42)
	machine.Memory[70] = uint8(0)

	for machine.Step() {

	}
}
