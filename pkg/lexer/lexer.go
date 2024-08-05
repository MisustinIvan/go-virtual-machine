package compiler

import "strings"

type Lexer struct {
	data []string
	pt   int
}

func New(input string) Lexer {
	return Lexer{
		data: strings.Fields(input),
		pt:   0,
	}
}

func (l *Lexer) NextToken() (string, bool) {
	if l.pt >= len(l.data) {
		return "", false
	} else {
		l.pt += 1
		return l.data[l.pt-1], true
	}

}

func (l *Lexer) ParseToken(input string) (Token, error) {
	if strings.HasSuffix(input, ":") {
		return Token{kind: LABEL_DEF, value: input}, nil
	} else {
		return Token{kind: UNKNOWN, value: input}, nil
	}
}

func (l *Lexer) Lex() ([]Token, error) {
	var res []Token
	for {
		tk, ct := l.NextToken()
		if !ct {
			return res, nil
		}

		pt, err := l.ParseToken(tk)
		if err != nil {
			return res, err
		} else {
			res = append(res, pt)
		}
	}
}
