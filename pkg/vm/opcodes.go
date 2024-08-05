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

func Str2Opcode(a string) opcode {
	switch a {
	case "NOP":
		return NOP
	case "ADD":
		return ADD
	case "ADI":
		return ADI
	case "SUB":
		return SUB
	case "SBI":
		return SBI
	case "MUL":
		return MUL
	case "MLI":
		return MLI
	case "DIV":
		return DIV
	case "DVI":
		return DVI
	case "MOV":
		return MOV
	case "MVI":
		return MVI
	case "PUS":
		return PUS
	case "POP":
		return POP
	case "JMP":
		return JMP
	case "JPI":
		return JPI
	case "INC":
		return INC
	case "DEC":
		return DEC
	case "LOD":
		return LOD
	case "LD8":
		return LD8
	case "STR":
		return STR
	case "ST8":
		return ST8
	case "PRT":
		return PRT
	case "PR8":
		return PR8
	case "HLT":
		return HLT
	case "CMP":
		return CMP
	case "JIF":
		return JIF
	case "JID":
		return JID
	case "JIZ":
		return JIZ
	case "JIE":
		return JIE
	case "JNE":
		return JNE
	case "JIG":
		return JIG
	case "JIS":
		return JIS
	case "JEG":
		return JEG
	case "JES":
		return JES
	case "XOR":
		return XOR
	}
	return NOP
}
