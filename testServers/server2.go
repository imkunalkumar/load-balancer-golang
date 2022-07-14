package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello from server2\n")
}

func main() {

	http.HandleFunc("/", hello)
	fmt.Println("server2 running on port :8082...")
	http.ListenAndServe(":8082", nil)
}
