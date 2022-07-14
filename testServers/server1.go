package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello from server1\n")
}

func main() {

	http.HandleFunc("/", hello)
	fmt.Println("server1 running on port :8081...")
	http.ListenAndServe(":8081", nil)

}
