package compiler

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Ghibranalj/jago/jago_compiler/utils"
)

const (
	progName = "program"
)

//Compile returns path to compiled path

func Compile(code string) (string, error) {

	cat := categorize(code)
	_ = cat
	fileContent := complete(code, cat)

	err := utils.WriteToFile(progName, fileContent)
	if err != nil {
		return "", err
	}

	err = javac("./" + progName + ".java")
	if err != nil {
		return "", err
	}

	return progName, nil
}

// TODO make a proper test
func javac(javaFile string) error {
	cmd := exec.Command("javac", javaFile)
	var errbuf strings.Builder
	cmd.Stderr = &errbuf
	cmd.Run()

	extCode := cmd.ProcessState.ExitCode()
	if extCode != 0 {
		return fmt.Errorf(errbuf.String())
	}
	return nil
}
