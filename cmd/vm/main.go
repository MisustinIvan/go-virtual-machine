package main

import (
	"go-virtual-machine/pkg/vm"
)

func main() {
	machine := vm.New(true)

	// LD8 RA 69
	machine.Memory[0] = uint8(vm.LD8)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(69)
	machine.Memory[3] = uint8(0)

	// PRT RA
	machine.Memory[4] = uint8(vm.PRT)
	machine.Memory[5] = uint8(vm.RA)

	// LD8 RB 70
	machine.Memory[6] = uint8(vm.LD8)
	machine.Memory[7] = uint8(vm.RB)
	machine.Memory[8] = uint8(70)
	machine.Memory[9] = uint8(0)

	// PRT RB
	machine.Memory[10] = uint8(vm.PRT)
	machine.Memory[11] = uint8(vm.RB)

	// ADD RA RB
	machine.Memory[12] = uint8(vm.ADD)
	machine.Memory[13] = uint8(vm.RA)
	machine.Memory[14] = uint8(vm.RB)

	// PRT RA
	machine.Memory[15] = uint8(vm.PRT)
	machine.Memory[16] = uint8(vm.RA)

	// ST8 71 RA
	machine.Memory[17] = uint8(vm.ST8)
	machine.Memory[18] = uint8(71)
	machine.Memory[19] = uint8(0)
	machine.Memory[20] = uint8(vm.RA)

	// LD8 RC
	machine.Memory[21] = uint8(vm.LD8)
	machine.Memory[22] = uint8(vm.RC)
	machine.Memory[23] = uint8(71)
	machine.Memory[24] = uint8(0)

	// PRT RC
	machine.Memory[25] = uint8(vm.PRT)
	machine.Memory[26] = uint8(vm.RC)

	// HLT
	machine.Memory[27] = uint8(vm.HLT)

	// constant data
	machine.Memory[69] = uint8(69)
	machine.Memory[70] = uint8(42)

	for machine.Step() {

	}
}
