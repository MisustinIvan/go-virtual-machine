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

func (t Token) Kind() TokenKind {
	return t.kind
}

func (t Token) Value() string {
	return t.value
}
