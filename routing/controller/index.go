package controller

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write([]byte("Index"))
	} else {
		w.Write([]byte("Wrong Method"))
	}
}
