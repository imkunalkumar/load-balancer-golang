package main

import "time"

func getAccessTime(server *server, i int) int64 {
	return server.serverLog[i].accessedOn
}
func getStatusCode(server *server, i int) int {
	return int(server.serverLog[i].statusCode)
}

func updateAccess(server *server, sCode int) {
	currentTime := time.Now().Unix()
	if server.serverLog[0].accessedOn == 0 {
		server.serverLog[0].accessedOn = currentTime
		server.serverLog[0].statusCode = int64(sCode)
	} else if server.serverLog[1].accessedOn == 0 {
		server.serverLog[1].accessedOn = currentTime
		server.serverLog[1].statusCode = int64(sCode)
	} else if server.serverLog[2].accessedOn == 0 {
		server.serverLog[2].accessedOn = currentTime
		server.serverLog[2].statusCode = int64(sCode)
	} else {
		server.serverLog[0] = server.serverLog[1]
		server.serverLog[1] = server.serverLog[2]
		server.serverLog[2].accessedOn = currentTime
		server.serverLog[2].statusCode = int64(sCode)
	}

}
