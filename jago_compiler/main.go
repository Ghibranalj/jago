package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ghibranalj/jago/jago_compiler/compiler"
	"github.com/Ghibranalj/jago/jago_compiler/utils"
)

func main() {
	http.HandleFunc("/version", Version)
	http.HandleFunc("/code", Code)
	fmt.Println(http.ListenAndServe(":1234", nil))

}

func Version(w http.ResponseWriter, r *http.Request) {
	stdout := ""
	stderr := ""
	_, err := utils.Command("javac --version", &stdout, &stderr)
	if err != nil {
		fmt.Fprintf(w, "%s \n %s", stdout, stderr)
		return
	}
	fmt.Fprintf(w, "%s \n %s", stdout, stderr)
}

type input struct {
	Code string
	Args string
}

func Code(w http.ResponseWriter, r *http.Request) {

	var inp input

	err := json.NewDecoder(r.Body).Decode(&inp)
	if err != nil {
		fmt.Fprintf(w, "400 : %s", err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		return
	}
	code := inp.Code

	progName, err := compiler.Compile(code)

	if err != nil {
		fmt.Fprintf(w, "compile error: %s", err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := execute(progName, inp.Args)
	if err != nil {
		fmt.Fprintf(w, "runtime error: %s", err.Error())
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", out)
	w.WriteHeader(http.StatusOK)
}

func execute(path, args string) (string, error) {

	stdout := ""
	stderr := ""
	cmd := "java " + path + args
	ext, err := utils.Command(cmd, &stdout, &stderr)
	if err != nil || ext != 0 {
		return stdout + stderr, err
	}
	utils.RemoveFile(path + ".class")
	return stdout + stderr, nil
}
