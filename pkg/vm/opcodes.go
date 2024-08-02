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
	LD8
	STR
	ST8
	PRT
	HLT
	CMP
	JIF
	JID
	JIZ
	JIE
	JNE
	JIG
	JIS
	JEG
	JES
	XOR
)
