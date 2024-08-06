# Usage
 - for reference on the available instructions and their signature see [architecture](./ARCHITECTURE.md)
 - first compile the `vm` and the `compiler` using `go build -o vm cmd/vm/main.go` and `go build -o compiler cmd/compiler/main.go`
 - then compile the test program [test.asm](./test.asm) using `./compiler test.asm test`
 - then run the program on the vm using `./vm test` and see the digits of the fibonacci sequence
