package main

import "net/http"

func redirectRequest(res http.ResponseWriter, req *http.Request) {
	server, err := getHealthyServer()
	if err != nil {
		http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	statusCode := 0
	server.redirectUrl.ModifyResponse = func(r *http.Response) error {

		res := r.StatusCode

		if res != 200 {

			updateAccess(server, 503)
			return nil
		}
		statusCode = res
		updateAccess(server, statusCode)
		return nil
	}
	server.redirectUrl.ServeHTTP(res, req)
	if statusCode == 0 {
		updateAccess(server, 503)
	}
}
