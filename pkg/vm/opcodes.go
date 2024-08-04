package vm

type opcode uint8

const (
	NOP opcode = iota
	ADD
	ADI
	SUB
	SBI
	MUL
	MLI
	DIV
	DVI
	MOV
	MVI
	PUS
	POP
	JMP
	JPI
	INC
	DEC
	LOD
	LD8
	STR
	ST8
	PRT
	PR8
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
