package main

import (
	"fmt"
	"log"
	"net/http"
)

var servers = []*server{}

var lastVisitedServer = 0

func main() {
	http.HandleFunc("/proxy", redirectRequest)
	http.HandleFunc("/url/register", registerRoute)
	go passiveCheck()
	fmt.Println("loadbalancer server running on port :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
