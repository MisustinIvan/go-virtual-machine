package vm

type sp_code uint16

const (
	OVERFLOW sp_code = iota
	DIVZERO
	ZERO
	EQUALS
	NOT_EQUALS
	GREATER
	SMALLER
	GREATER_OR_EQUAL
	SMALLER_OR_EQUAL
)
