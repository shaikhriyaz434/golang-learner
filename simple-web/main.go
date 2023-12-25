package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func Handler(response http.ResponseWriter, request *http.Request) {
	content, err := os.Open("./main.go")
	if err != nil {
		fmt.Print("failed to read a file")
	}
	io.Copy(response, content)
}
