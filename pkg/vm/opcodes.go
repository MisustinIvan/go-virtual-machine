package vm

type opcode uint8

const (
	NOP opcode = iota
	ADD
	SUB
	MUL
	DIV
	MOV
	JMP
	INC
	DEC
	LOD
	STR
	PRT
	HLT
	CMP
	JIO
	JIZ
	JIE
	JNE
	JIG
	JIS
	JEG
	JES
	XOR
)
