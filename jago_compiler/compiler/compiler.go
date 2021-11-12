package compiler

import (
	"fmt"

	"github.com/Ghibranalj/jago/jago_compiler/utils"
)

const (
	progName = "program"
)

// Compile compiles code to  a .class file
//
// returns path to .class file path and compilation error if any
func Compile(code string) (string, error) {

	cat := categorize(code)
	_ = cat
	fileContent := complete(code, cat)

	javaFile := "./" + progName + ".java"
	err := utils.WriteToFile(javaFile, fileContent)
	if err != nil {
		return "", err
	}

	err = javac(javaFile)

	if err != nil {
		return "", err
	}

	err = utils.RemoveFile(javaFile)
	if err != nil {
		return "", err
	}

	return progName, nil
}

// TODO make a proper test
func javac(javaFile string) error {
	stdout := ""
	stderr := ""
	ext, err := utils.Command("javac "+javaFile, &stdout, &stderr)
	if err != nil || ext != 0 {
		return fmt.Errorf("%s \n %s", stdout, stderr)
	}
	return nil
}
