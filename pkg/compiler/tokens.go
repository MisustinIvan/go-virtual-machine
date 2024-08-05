package compiler

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

type Token struct {
	kind  TokenKind
	value string
}

func (t Token) Kind() TokenKind {
	return t.kind
}

func (t Token) Vind() string {
	return t.value
}
