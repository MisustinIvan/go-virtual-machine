package preprocessor_test

import (
	"go-virtual-machine/pkg/preprocessor"
	"strings"
	"testing"
)

func TestPreprocessor(t *testing.T) {
	var input string = "#begin LIGMA balls in your mouth #end LIGMA"
	var expected_output string = "balls in your mouth"
	input = strings.Join(strings.Fields(input), " ")

	output, err := preprocessor.Preprocess(input)
	if err != nil {
		t.Logf("INPUT: %s\n", input)
		t.Logf("OUTPUT: %s\n", output)
		t.Fatalf("Preprocessor failed: %s\n", err.Error())
	}

	if output != expected_output {
		t.Logf("OUTPUT: %s\n", output)
		t.Logf("EXPECTED_OUTPUT: %s\n", expected_output)
		t.Fatal("Preprocessor produced unexpected result...")
	}
}

func TestUnterminatedMacro(t *testing.T) {
	var input1 string = "#begin LIGMA balls in your mouth LIGMA"
	input1 = strings.Join(strings.Fields(input1), " ")

	output, err := preprocessor.Preprocess(input1)
	if err == nil {
		t.Logf("INPUT: %s\n", input1)
		t.Logf("OUTPUT: %s\n", output)
		t.Fatalf("Preprocessor succeeded, which is wrong\n")
	}

	var input2 string = "#begin"
	input2 = strings.Join(strings.Fields(input2), " ")

	output2, err := preprocessor.Preprocess(input2)
	if err == nil {
		t.Logf("INPUT: %s\n", input2)
		t.Logf("OUTPUT: %s\n", output2)
		t.Fatalf("Preprocessor succeeded, which is wrong\n")
	}
}
