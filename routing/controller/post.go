package controller

import "net/http"

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "POST" {
		r.ParseForm()
		input := r.Form.Get("input")
		w.Write([]byte(input))
	} else {
		w.Write([]byte("Wrong Method"))
	}
}
