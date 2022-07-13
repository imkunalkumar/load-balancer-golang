package main

import (
	"log"
	"net/http"
)

var servers = []*server{}

var lastVisitedServer = 0

func main() {
	http.HandleFunc("/proxy", redirectRequest)
	http.HandleFunc("/url/register", registerRoute)
	go passiveCheck()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
