package lexer

type keywords []string

func (w keywords) contains(el string) bool {
	for _, value := range w {
		if value == el {
			return true
		}
	}
	return false
}

var keywords_regs keywords = []string{
	"PC",
	"BP",
	"SP",
	"RS",
	"RA",
	"RB",
	"RC",
	"RD",
}

const ENTRY string = "_start"

var keywords_inst keywords = []string{
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
