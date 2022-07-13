package main

import "net/http/httputil"

type recievedUrl struct {
	ServerName string `json:"ServerName"`
	Url        string `json:"url"`
}

type server struct {
	Name        string
	URL         string
	redirectUrl *httputil.ReverseProxy
	serverLog   [3]accessLog
}
type accessLog struct {
	accessedOn int64
	statusCode int64
}
