package compiler

var keywords_regs []string = []string{
	"PC",
	"BP",
	"SP",
	"RS",
	"RA",
	"RB",
	"RC",
	"RD",
}

var keywords_inst []string = []string{
	"CAL",
	"RET",
	// the two extra ones
	"NOP",
	"ADD",
	"ADI",
	"SUB",
	"SBI",
	"MUL",
	"MLI",
	"DIV",
	"DVI",
	"MOV",
	"MVI",
	"PUS",
	"POP",
	"JMP",
	"JPI",
	"INC",
	"DEC",
	"LOD",
	"LD8",
	"STR",
	"ST8",
	"PRT",
	"PR8",
	"HLT",
	"CMP",
	"JIF",
	"JID",
	"JIZ",
	"JIE",
	"JNE",
	"JIG",
	"JIS",
	"JEG",
	"JES",
	"XOR",
}
