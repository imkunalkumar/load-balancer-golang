package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func CreatenewServer(name, urlStr string) *server {
	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	accessObj := [3]accessLog{}
	return &server{
		Name:        name,
		URL:         urlStr,
		redirectUrl: rp,
		serverLog:   accessObj,
	}
}

func registerRoute(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		resurl := recievedUrl{}
		body, _ := io.ReadAll(req.Body)
		json.Unmarshal(body, &resurl)
		serverInstance := CreatenewServer(resurl.ServerName, resurl.Url)
		servers = append(servers, serverInstance)

	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
