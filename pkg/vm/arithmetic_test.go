package vm_test

import (
	"go-virtual-machine/pkg/vm"
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)
		machine.SetReg(vm.RB, n1)

		machine.Memory[0] = uint8(vm.ADD)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(vm.RB)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		if int(machine.GetReg(vm.RA)) != int(n0)+int(n1) && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestAdi(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)

		machine.Memory[0] = uint8(vm.ADI)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(n1)
		machine.Memory[3] = uint8(n1 >> 8)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		if int(machine.GetReg(vm.RA)) != int(n0)+int(n1) && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Logf("N0: 0x%x\n", n0)
			t.Logf("N1: 0x%x\n", n1)
			t.Logf("RA: 0x%x\n", machine.GetReg(vm.RA))
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestSub(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)
		machine.SetReg(vm.RB, n1)

		machine.Memory[0] = uint8(vm.SUB)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(vm.RB)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		if int(machine.GetReg(vm.RA)) != int(n0)-int(n1) && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Logf("N0 %d\n", n0)
			t.Logf("N1 %d\n", n1)
			t.Logf("RA %d\n", machine.GetReg(vm.RA))
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestSbi(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)

		machine.Memory[0] = uint8(vm.SBI)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(n1)
		machine.Memory[3] = uint8(n1 >> 8)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		if int(machine.GetReg(vm.RA)) != int(n0)-int(n1) && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Logf("N0 %d\n", n0)
			t.Logf("N1 %d\n", n1)
			t.Logf("RA %d\n", machine.GetReg(vm.RA))
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestMul(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)
		machine.SetReg(vm.RB, n1)

		machine.Memory[0] = uint8(vm.MUL)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(vm.RB)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		if int(machine.GetReg(vm.RA)) != int(n0)*int(n1) && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Logf("N0 %d\n", n0)
			t.Logf("N1 %d\n", n1)
			t.Logf("RA %d\n", machine.GetReg(vm.RA))
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestMli(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)

		machine.Memory[0] = uint8(vm.MLI)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(n1)
		machine.Memory[3] = uint8(n1 >> 8)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		if int(machine.GetReg(vm.RA)) != int(n0)*int(n1) && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Logf("N0 %d\n", n0)
			t.Logf("N1 %d\n", n1)
			t.Logf("RA %d\n", machine.GetReg(vm.RA))
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestDiv(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		if n1 == 0 {
			continue
		}

		machine.SetReg(vm.RA, n0)
		machine.SetReg(vm.RB, n1)

		machine.Memory[0] = uint8(vm.DIV)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(vm.RB)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		t.Logf("N0 %d\n", n0)
		t.Logf("N1 %d\n", n1)
		t.Logf("RA %d\n", machine.GetReg(vm.RA))
		t.Logf("RES %d\n", int(n0)/int(n1))

		if int(machine.GetReg(vm.RA)) != int(n0)/int(n1) {
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestDvi(t *testing.T) {
	var n0 uint16
	var n1 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())
		n1 = uint16(rand.Uint32())

		if n1 == 0 {
			continue
		}

		machine.SetReg(vm.RA, n0)

		machine.Memory[0] = uint8(vm.DVI)
		machine.Memory[1] = uint8(vm.RA)
		machine.Memory[2] = uint8(n1)
		machine.Memory[3] = uint8(n1 >> 8)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		if int(machine.GetReg(vm.RA)) != int(n0)/int(n1) {
			t.Logf("N0 %d\n", n0)
			t.Logf("N1 %d\n", n1)
			t.Logf("RA %d\n", machine.GetReg(vm.RA))
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestInc(t *testing.T) {
	var n0 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)

		machine.Memory[0] = uint8(vm.INC)
		machine.Memory[1] = uint8(vm.RA)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		t.Logf("N0 %d\n", n0)

		if int(machine.GetReg(vm.RA)) != int(n0)+1 && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Fatal("instruction returned wrong result")
		}
	}
}

func TestDec(t *testing.T) {
	var n0 uint16

	for i := 0; i < 20; i++ {
		machine := vm.New(false)

		n0 = uint16(rand.Uint32())

		machine.SetReg(vm.RA, n0)

		machine.Memory[0] = uint8(vm.DEC)
		machine.Memory[1] = uint8(vm.RA)

		if !machine.Step() {
			t.Fatal("instruction caused halt")
		}

		t.Logf("N0 %d\n", n0)

		if int(machine.GetReg(vm.RA)) != int(n0)-1 && machine.GetReg(vm.RS) != uint16(vm.OVERFLOW) {
			t.Fatal("instruction returned wrong result")
		}
	}
}
