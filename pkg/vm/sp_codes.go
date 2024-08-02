package vm

type sp_code uint16

const (
	OVERFLOW sp_code = iota
	DIVZERO
	ZERO
	EQUALS
	GREATER
	SMALLER
)
