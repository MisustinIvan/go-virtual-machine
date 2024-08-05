package main

import (
	"fmt"
	"go-virtual-machine/pkg/compiler"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Provide a source code file and a output file name maggot")
		os.Exit(1)
	}

	err := compiler.Compile(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
}
