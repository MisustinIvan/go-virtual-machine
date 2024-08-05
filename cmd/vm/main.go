package main

import (
	"fmt"
	"go-virtual-machine/pkg/vm"
	"io"
	"os"
	"path"
)

func main() {
	machine := vm.New(false)

	if len(os.Args) < 2 {
		fmt.Println("Please provide binary file")
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fp, err := os.Open(path.Join(cwd, os.Args[1]))
	if err != nil {
		panic(err)
	}

	program, err := io.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	fp.Close()

	for i, b := range program {
		machine.Memory[i] = b
	}

	for machine.Step() {
	}
}
