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

func TestJif(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 42)
	machine.SetReg(vm.RB, 69)

	machine.Memory[0] = uint8(vm.SUB)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JIF)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}

func TestJid(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 69)
	machine.SetReg(vm.RB, 0)

	machine.Memory[0] = uint8(vm.DIV)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JID)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}

func TestJiz(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 69)
	machine.SetReg(vm.RB, 69)

	machine.Memory[0] = uint8(vm.SUB)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JIZ)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}
func TestJie(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 69)
	machine.SetReg(vm.RB, 69)

	machine.Memory[0] = uint8(vm.CMP)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JIE)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}
func TestJne(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 69)
	machine.SetReg(vm.RB, 42)

	machine.Memory[0] = uint8(vm.CMP)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JNE)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}
func TestJig(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 69)
	machine.SetReg(vm.RB, 42)

	machine.Memory[0] = uint8(vm.CMP)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JIG)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}
func TestJis(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 42)
	machine.SetReg(vm.RB, 69)

	machine.Memory[0] = uint8(vm.CMP)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JIS)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}
func TestJeg(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 69)
	machine.SetReg(vm.RB, 42)

	machine.Memory[0] = uint8(vm.CMP)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JEG)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}
func TestJes(t *testing.T) {
	machine := vm.New(false)

	machine.SetReg(vm.RA, 42)
	machine.SetReg(vm.RB, 69)

	machine.Memory[0] = uint8(vm.CMP)
	machine.Memory[1] = uint8(vm.RA)
	machine.Memory[2] = uint8(vm.RB)

	machine.Memory[3] = uint8(vm.JES)
	machine.Memory[4] = uint8(69)
	machine.Memory[5] = uint8(0)

	if !machine.Step() {
		t.Fatal("instruction caused halt")
	} else if !machine.Step() {
		t.Fatal("instruction caused halt")
	}

	if machine.GetReg(vm.PC) != 69 {
		t.Logf("PC: %d", machine.GetReg(vm.PC))
		t.Logf("SP: %d", machine.GetReg(vm.SP))
		t.Fatal("wrong program counter")
	}
}
