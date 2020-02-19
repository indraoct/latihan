package controller

import (
	"encoding/json"
	"latihan/routing/lib/request"
	"latihan/routing/lib/response"
	"net/http"
)

func PostJsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "POST" {
		var input request.InputJson
		var output response.OutputJson
		var outputJson []byte
		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			output.Status = 0
			output.Message = err.Error()
			outputJson, _ = json.Marshal(output)
			w.WriteHeader(500)
			w.Write([]byte(outputJson))
			return
		}

		output.Status = 1
		output.Message = "Success"
		outputJson, _ = json.Marshal(output)
		w.WriteHeader(200)
		w.Write([]byte(outputJson))

	} else {
		w.WriteHeader(500)
		w.Write([]byte("Wrong Method"))
	}
}
