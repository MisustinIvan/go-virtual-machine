package compiler

import (
	"fmt"
	"go-virtual-machine/pkg/lexer"
	"go-virtual-machine/pkg/parser"
	"go-virtual-machine/pkg/preprocessor"
	"go-virtual-machine/pkg/translator"
	"io"
	"os"
	"path"
)

func Compile(input_name string, output_name string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fp, err := os.Open(path.Join(cwd, input_name))
	if err != nil {
		return err
	}
	defer fp.Close()

	fmt.Printf("Compiling %s\n", path.Join(cwd, input_name))

	fc, err := io.ReadAll(fp)
	if err != nil {
		return err
	}

	file := string(fc)

	preprocessed, err := preprocessor.Preprocess(file)
	if err != nil {
		return err
	}

	lexed, err := lexer.Lex(preprocessed)
	if err != nil {
		return err
	}

	parsed, err := parser.Parse(lexed)
	if err != nil {
		return err
	}

	fmt.Println(parsed)

	translated, err := translator.Translate(parsed)
	if err != nil {
		return err
	}

	dp, err := os.Create(path.Join(cwd, output_name))
	if err != nil {
		return err
	}

	_, err = dp.Write(translated)
	if err != nil {
		return err
	}

	fmt.Println("File written to disk")

	return nil
}
