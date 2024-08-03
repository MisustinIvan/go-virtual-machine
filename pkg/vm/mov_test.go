package vm_test

import (
	"go-virtual-machine/pkg/vm"
	"math/rand"
	"testing"
)

func TestMov(t *testing.T) {
	machine := vm.New(false)

	var val uint16 = 16

	machine.SetReg(vm.RA, val)

	machine.Memory[0] = uint8(vm.MOV)
	machine.Memory[1] = uint8(vm.RB)
	machine.Memory[2] = uint8(vm.RA)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.RA) != machine.GetReg(vm.RB) || machine.GetReg(vm.RA) != val {
		t.Fatal("mov instruction failed")
	}
}

func TestLod(t *testing.T) {
	var address uint16
	var number uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)
		address = uint16(rand.Int31())
		number = uint16(rand.Int31())

		if address <= 3 || address >= vm.MEMSIZE-2 {
			continue
		}

		machine.Memory[0] = uint8(vm.LOD)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(address)
		machine.Memory[3] = uint8(address >> 8)

		machine.Memory[address] = uint8(number)
		machine.Memory[address+1] = uint8(number >> 8)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		lower := uint8(machine.GetReg(vm.RA))
		higher := uint8(machine.GetReg(vm.RA) >> 8)

		val := uint16(lower) | (uint16(higher) << 8)

		t.Logf("Lower: 0x%x\n", lower)
		t.Logf("Higher: 0x%x\n", higher)
		t.Logf("Val: 0x%x\n", val)
		t.Logf("Address: 0x%x\n", address)
		t.Logf("Number: 0x%x\n", number)

		if val != number {
			t.Fatal("Loaded value not same as number")
		}
	}
}

func TestLod8(t *testing.T) {
	var address uint16
	var number uint8

	for i := 0; i < 20; i++ {
		machine := vm.New(false)
		address = uint16(rand.Int31())
		number = uint8(rand.Int31())

		if address <= 3 || address >= vm.MEMSIZE-2 {
			continue
		}

		machine.Memory[0] = uint8(vm.LD8)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(address)
		machine.Memory[3] = uint8(address >> 8)

		machine.Memory[address] = number

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		lower := uint8(machine.GetReg(vm.RA))

		t.Logf("Lower: 0x%x\n", lower)
		t.Logf("Address: 0x%x\n", address)
		t.Logf("Number: 0x%x\n", number)

		if lower != number {
			t.Fatal("Loaded value not same as number")
		}
	}
}

func TestStr(t *testing.T) {
	var number uint16
	var address uint16

	for i := 0; i < 20; i++ {
		number = uint16(rand.Int())
		address = uint16(rand.Int())

		machine := vm.New(false)

		if address <= 3 || address >= vm.MEMSIZE-2 {
			continue
		}

		machine.SetReg(vm.RA, number)

		machine.Memory[0] = uint8(vm.STR)
		machine.Memory[1] = uint8(address)
		machine.Memory[2] = uint8(address >> 8)
		machine.Memory[3] = uint8(vm.RA)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		lower := machine.Memory[address]
		higher := machine.Memory[address+1]

		t.Logf("Lower: 0x%x\n", lower)
		t.Logf("higher: 0x%x\n", higher)

		val := uint16(lower) | (uint16(higher) << 8)
		t.Logf("Val: 0x%x\n", val)

		if val != number {
			t.Fatal("Value stored not original")
		}
	}
}

func TestStr8(t *testing.T) {
	var number uint8

	for i := 0; i < 20; i++ {
		number = uint8(rand.Int())
		address := uint16(rand.Int())

		if address <= 3 || address >= vm.MEMSIZE-2 {
			continue
		}

		machine := vm.New(false)

		machine.SetReg(vm.RA, uint16(number))

		machine.Memory[0] = uint8(vm.STR)
		machine.Memory[1] = uint8(address)
		machine.Memory[2] = uint8(address >> 8)
		machine.Memory[3] = uint8(vm.RA)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		lower := machine.Memory[address]

		t.Logf("Lower: 0x%x\n", lower)

		if lower != number {
			t.Fatal("Value stored not original")
		}
	}
}
