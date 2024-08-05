package compiler

import (
	"fmt"
	"go-virtual-machine/pkg/preprocessor"
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

	fmt.Println(preprocessed)

	return nil
}
