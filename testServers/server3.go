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
	fmt.Println("server3 running on port :8083...")
	http.ListenAndServe(":8083", nil)
}
