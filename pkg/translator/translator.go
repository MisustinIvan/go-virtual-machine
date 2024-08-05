package translator

import (
	"fmt"
	"go-virtual-machine/pkg/lexer"
	"go-virtual-machine/pkg/vm"
	"strconv"
)

func translate_instruction(tk lexer.Token) byte {
	return byte(vm.Str2Opcode(tk.Value()))
}

func translate_dec_literal(tk lexer.Token) (byte, byte) {
	res, _ := strconv.Atoi(tk.Value())
	return byte(res), byte(res >> 8)
}

func translate_hex_literal(tk lexer.Token) (byte, byte) {
	res, _ := strconv.ParseUint(tk.Value()[2:], 16, 16)
	return byte(res), byte(res >> 8)
}

func Translate(tokens []lexer.Token) ([]byte, error) {
	result := []byte{}
	for _, tk := range tokens {
		switch tk.Kind() {

		case lexer.INSTRUCTION:
			result = append(result, translate_instruction(tk))

		case lexer.DEC_LITERAL:
			lower, higher := translate_dec_literal(tk)
			result = append(result, lower)
			result = append(result, higher)

		case lexer.HEX_LITERAL:
			lower, higher := translate_hex_literal(tk)
			result = append(result, lower)
			result = append(result, higher)

		case lexer.REGISTER:
			result = append(result, byte(vm.Str2Reg(tk.Value())))

		case lexer.UNKNOWN:
			return result, fmt.Errorf("Encountered unknown token %s:%s", tk.Kind().String(), tk.Value())

		default:
			return result, fmt.Errorf("Encountered unknown token %s:%s", tk.Kind().String(), tk.Value())
		}
	}
	return result, nil
}
