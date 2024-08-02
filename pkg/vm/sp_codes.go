package vm

type sp_code uint16

const (
	NONE sp_code = iota
	OVERFLOW
	DIVZERO
	ZERO
	EQUALS
	GREATER
	SMALLER
)
