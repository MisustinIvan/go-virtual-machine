package lexer

type TokenKind int

const (
	UNKNOWN TokenKind = iota
	DEC_LITERAL
	HEX_LITERAL
	INSTRUCTION
	REGISTER
	LABEL_DEF
	LABEL_REF
)

func (t TokenKind) String() string {
	switch t {
	case DEC_LITERAL:
		return "DEC_LITERAL"
	case HEX_LITERAL:
		return "HEX_LITERAL"
	case INSTRUCTION:
		return "INSTRUCTION"
	case REGISTER:
		return "REGISTER"
	case LABEL_DEF:
		return "LABEL_DEF"
	case LABEL_REF:
		return "LABEL_REF"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	kind  TokenKind
	value string
}

// used for use outside of lexer package... because i am not doing uppercase member fields
func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind:  kind,
		value: value,
	}
}

func (t Token) Kind() TokenKind {
	return t.kind
}

func (t Token) Value() string {
	return t.value
}

// size of the token in code (bytes)
func (t Token) Size() int {
	switch t.Kind() {
	case DEC_LITERAL:
		return 2
	case HEX_LITERAL:
		return 2
	case INSTRUCTION:
		// those two expand to some more ass-embly code down the line
		if t.value == "CAL" {
			return 15
		} else if t.value == "RET" {
			return 6
		} else {
			return 1
		}
	case REGISTER:
		return 1
	case LABEL_DEF:
		return 0
	case LABEL_REF:
		return 2
	default:
		return 0
	}
}
