package lexer

import (
	"fmt"
	"strings"
)

type Lexer struct {
	data   []string
	labels keywords
	result []Token
	tp     int
}

func New(input string) Lexer {
	return Lexer{
		data:   strings.Fields(input),
		result: []Token{},
		tp:     0,
	}
}

func (l *Lexer) next_token() (string, bool) {
	if l.tp >= len(l.data) {
		return "", false
	} else {
		l.tp += 1
		return l.data[l.tp-1], true
	}

}

func isNumberDec(input string) bool {
	for _, c := range input {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func isNumberHex(input string) bool {
	for _, c := range input {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}
	return true
}

func ReprTokens(arg []Token) string {
	res := "["
	for _, tk := range arg {
		res += "{"
		res += tk.Kind().String()
		res += " "
		res += tk.Value()
		res += "}"
	}

	res += "]"

	return res
}

func (l *Lexer) token_kind(input string) TokenKind {
	if keywords_inst.contains(input) {
		return INSTRUCTION
	} else if keywords_regs.contains(input) {
		return REGISTER
	} else if strings.HasPrefix(input, ":") {
		return LABEL_DEF
	} else if l.labels.contains(input) {
		return LABEL_REF
	} else if isNumberDec(input) {
		return DEC_LITERAL
	} else if strings.HasPrefix(input, "0x") && isNumberHex(input[2:]) {
		return HEX_LITERAL
	} else {
		return UNKNOWN
	}
}

func Lex(input string) ([]Token, error) {
	l := New(input)

	for {
		tk, ok := l.next_token()
		if !ok {
			return l.result, nil
		}

		switch kind := l.token_kind(tk); kind {
		case UNKNOWN:
			l.result = append(l.result, Token{
				kind:  kind,
				value: tk,
			})
			return l.result, fmt.Errorf("Unknown token detected: %s", tk)
		case LABEL_DEF:
			l.labels = append(l.labels, tk[1:])
			l.result = append(l.result, Token{
				kind:  kind,
				value: tk,
			})
			break
		default:
			l.result = append(l.result, Token{
				kind:  kind,
				value: tk,
			})
			break
		}
	}
}
