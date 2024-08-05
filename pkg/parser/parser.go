package parser

import (
	"fmt"
	"go-virtual-machine/pkg/lexer"
)

type Parser struct {
	ip     int
	tokens []lexer.Token
	// maps label string to address
	labels map[string]uint16
	result []lexer.Token
}

func New(tokens []lexer.Token) (Parser, error) {
	labels := map[string]uint16{}
	addr := uint16(0)

	for _, t := range tokens {
		if t.Kind() == lexer.LABEL_DEF {
			labels[t.Value()] = addr
		}

		if int(addr)+t.Size() > 65535 {
			return Parser{
				ip:     0,
				tokens: tokens,
				labels: labels,
				result: []lexer.Token{},
			}, fmt.Errorf("The machine can't have more RAM you maggot")
		}

		addr += uint16(t.Size())
	}

	res := []lexer.Token{}
	start_addr, _ := labels["_start"]
	if start_addr != 0 {
		res = append(res, lexer.NewToken(lexer.INSTRUCTION, "JMP"))
		res = append(res, lexer.NewToken(lexer.DEC_LITERAL, fmt.Sprintf("%d", labels["_start"])))
	}

	return Parser{
		ip:     0,
		tokens: tokens,
		labels: labels,
		result: res,
	}, nil
}

func (p *Parser) GetNextToken() (lexer.Token, bool) {
	if p.ip >= len(p.tokens) {
		return lexer.NewToken(lexer.UNKNOWN, ""), false
	}
	p.ip += 1
	return p.tokens[p.ip-1], true
}

func (p *Parser) Parse() (bool, error) {
	tk, ok := p.GetNextToken()
	if !ok {
		return false, nil
	}

	switch tk.Kind() {

	// this one aint regular...
	// but i mean why the hell not
	// it's quite simple...
	case lexer.INSTRUCTION:
		if tk.Value() != "CAL" && tk.Value() != "RET" {
			p.result = append(p.result, tk)
		} else {
			if tk.Value() == "CAL" {
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "PUS"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "BP"))
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "MOV"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "BP"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "SP"))
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "MOV"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "RD"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "PC"))
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "ADI"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "RD"))
				p.result = append(p.result, lexer.NewToken(lexer.DEC_LITERAL, "12"))
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "PUS"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "RD"))
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "JMP"))
			} else {
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "POP"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "RD"))
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "JPI"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "RD"))
				p.result = append(p.result, lexer.NewToken(lexer.INSTRUCTION, "POP"))
				p.result = append(p.result, lexer.NewToken(lexer.REGISTER, "BP"))
			}
		}

	case lexer.LABEL_REF:
		addr, ok := p.labels[tk.Value()]
		if !ok {
			return false, fmt.Errorf("Unknown label referenced %s at %d", tk.Value(), p.ip)
		} else {
			p.result = append(p.result, lexer.NewToken(lexer.DEC_LITERAL, fmt.Sprintf("%d", addr)))
		}

	case lexer.LABEL_DEF:

	case lexer.UNKNOWN:
		return false, fmt.Errorf("Unknown token %s:%s at %d", tk.Value(), tk.Kind().String(), p.ip)

	default:
		p.result = append(p.result, tk)
	}

	return true, nil
}

func Parse(tokens []lexer.Token) ([]lexer.Token, error) {
	p, e := New(tokens)
	if e != nil {
		return p.result, e
	}

	var err error = nil

	for {
		cont, e := p.Parse()
		if e != nil {
			err = e
			break
		} else if !cont {
			break
		}
	}

	return p.result, err
}
