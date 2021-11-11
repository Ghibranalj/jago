package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println(http.ListenAndServe(":1234", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("javac", "--version")
	out, err := cmd.Output()
	if err != nil {
		out = []byte(err.Error())
	}
	fmt.Fprintf(w, "%s", string(out))
}
