package preprocessor

import (
	"fmt"
	"strings"
)

type preprocessor struct {
	data   []string
	macros map[string][]string
	result []string
	dp     int
}

func (p *preprocessor) next_token() (string, bool) {
	if p.dp >= len(p.data) {
		return "", false
	}
	p.dp += 1
	return p.data[p.dp-1], true
}

func (p *preprocessor) new_macro(name string) {
	p.macros[name] = []string{}
}

func (p *preprocessor) macro_append(name string, value string) {
	p.macros[name] = append(p.macros[name], value)
}

func new(a string) preprocessor {
	return preprocessor{
		data:   strings.Fields(a),
		macros: map[string][]string{},
		result: []string{},
		dp:     0,
	}
}

func Preprocess(input string) (string, error) {
	p := new(input)

	for {
		tk, ok := p.next_token()
		if !ok {
			break
		}

		if tk == "#begin" {
			name, ok := p.next_token()
			if !ok {
				return strings.Join(p.result, " "), fmt.Errorf("unterminated macro directive")
			}
			p.new_macro(name)

			for {
				next_token, ok := p.next_token()
				if !ok {
					return strings.Join(p.result, " "), fmt.Errorf("unterminated macro directive")
				}

				if next_token == "#end" {
					break
				} else {
					p.macro_append(name, next_token)
				}
			}
		} else {
			macro, exists := p.macros[tk]
			if !exists {
				p.result = append(p.result, tk)
			} else {
				for _, mt := range macro {
					p.result = append(p.result, mt)
				}
			}
		}
	}

	return strings.Join(p.result, " "), nil
}
