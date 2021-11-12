package main

import (
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

func Code(w http.ResponseWriter, r *http.Request) {

	// get code from request
	// compile code
	// 		if error return error message
	// run program
	//		if error return error message
	// return output from program

	compiler.Compile("System.out.println()")
}

func execute(path string) (string, error) {

	return "", nil
}
