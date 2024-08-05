package compiler_test

import (
	"go-virtual-machine/pkg/compiler"
	"testing"
)

func TestCompiler(t *testing.T) {
	err := compiler.Compile("./test.asm", "test")

	if err != nil {
		t.Fatalf("Compilation failed: %s...\n", err.Error())
	}
}
