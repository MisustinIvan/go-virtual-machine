package lexer_test

import (
	"go-virtual-machine/pkg/lexer"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "MOV RA RB :LIGMA PUS RA JMP LIGMA LOD RC 0x45"
	expected_output := "[{INSTRUCTION MOV}{REGISTER RA}{REGISTER RB}{LABEL_DEF :LIGMA}{INSTRUCTION PUS}{REGISTER RA}{INSTRUCTION JMP}{LABEL_REF LIGMA}{INSTRUCTION LOD}{REGISTER RC}{HEX_LITERAL 0x45}]"

	res, err := lexer.Lex(input)
	if err != nil {
		t.Fatalf("Lexer failed with error: %s\n", err.Error())
	}

	if expected_output != lexer.ReprTokens(res) {
		t.Logf("OUTPUT: %s\n", lexer.ReprTokens(res))
		t.Logf("EXPECTED_OUTPUT: %s\n", expected_output)
		t.Fatal("Lexer produced unexpected result...")
	}
}
