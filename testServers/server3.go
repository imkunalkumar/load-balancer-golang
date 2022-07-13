package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello from server3\n")
}

func main() {

	http.HandleFunc("/", hello)

	http.ListenAndServe(":8083", nil)
}
