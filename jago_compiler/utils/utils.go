package utils

import (
	"os"
	"os/exec"
	"strings"
)

func WriteToFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}

// Command run a command standard
//
// standard output is written to stdout. standard error is written to stderr
//
// returns exitcode and error if any
func Command(cmdstr string, stdout, stderr *string) (int, error) {

	cmds := strings.Split(cmdstr, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	var outbuf, errbuf strings.Builder
	cmd.Stderr = &errbuf
	cmd.Stdout = &outbuf
	err := cmd.Run()

	*stdout = outbuf.String()
	*stderr = errbuf.String()

	if err != nil {
		return -1, err
	}

	extCode := cmd.ProcessState.ExitCode()

	if extCode != 0 {
		return extCode, err
	}

	return extCode, err
}
