package main

import (
	"main/pkg/vm"
)

func main() {
	machine := vm.New()

	machine.Memory[0] = uint8(vm.LOD)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(69)
	machine.Memory[3] = uint8(0)
	machine.Memory[4] = uint8(vm.MOV)
	machine.Memory[5] = uint8(vm.RB)
	machine.Memory[6] = uint8(vm.RA)
	machine.Memory[7] = uint8(vm.PRT)
	machine.Memory[8] = uint8(vm.RA)
	machine.Memory[9] = uint8(vm.PRT)
	machine.Memory[10] = uint8(vm.RB)
	machine.Memory[11] = uint8(vm.HLT)

	machine.Memory[69] = uint8(42)
	machine.Memory[70] = uint8(0)

	for machine.Step() {

	}
}
