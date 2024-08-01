package main

import (
	"main/pkg/vm"
)

func main() {
	machine, err := vm.New(64)
	if err != nil {
		panic(err)
	}

	machine.Memory[0] = uint8(vm.JMP)
	machine.Memory[1] = uint8(vm.MOV)
	machine.Memory[2] = uint8(vm.CMP)
	machine.Memory[3] = uint8(vm.NOP)
	machine.Memory[4] = uint8(vm.HLT)

	for machine.Step() {

	}
}
