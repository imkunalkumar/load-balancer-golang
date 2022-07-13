package main

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/robfig/cron.v2"
)

func getHealthyServer() (*server, error) {
	for i := 0; i < len(servers); i++ {
		server := getServer()

		if healthCheck(server) {
			return server, nil
		}

	}
	return nil, fmt.Errorf("no healthy hosts")
}

//round robin
func getServer() *server {
	if lastVisitedServer > len(servers)-1 {
		lastVisitedServer = 0
	}
	server := servers[lastVisitedServer]
	lastVisitedServer++
	return server
}

func healthCheck(server *server) bool {

	if getAccessTime(server, 0) == 0 {

		if initialHealthCheck(server) != 200 {
			return false
		} else {
			return true
		}
	}
	currentTime := time.Now().Unix()

	if currentTime-getAccessTime(server, 0) <= 15 && getStatusCode(server, 0) == 200 {
		if getAccessTime(server, 1) == 0 {

			return true
		}
		if getStatusCode(server, 1) == 200 && (getAccessTime(server, 2) == 0 || getStatusCode(server, 2) == 200) {

			return true
		}
	}
	return false
}

//it checks initial helth of endpoint
func initialHealthCheck(server *server) int {
	resp, err := http.Head(server.URL)
	if err != nil {
		updateAccess(server, 503)
		return 0
	}

	if resp.StatusCode != http.StatusOK {
		updateAccess(server, 503)
		return 0
	}
	updateAccess(server, 200)
	return http.StatusOK
}

//passive health check

func passiveCheck() {
	s := cron.New()
	s.AddFunc("@every 4s", func() {
		if len(servers) > 0 {
			for _, val := range servers {
				initialHealthCheck(val)
			}
		}

	})
	s.Start()
}
