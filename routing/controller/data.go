package controller

import (
	"database/sql"
	"encoding/json"
	"latihan/routing/lib/response"
	"latihan/routing/models"
	"net/http"
)

func GetDataHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		var bahasa models.Bahasa
		var arr_bahasa []models.Bahasa
		var output response.OutputDataJson

		rows, err := db.Query("SELECT id,name FROM bahasa ORDER BY id desc")
		if err != nil {

			output.Status = 0
			output.Message = err.Error()
			outputJson, _ := json.Marshal(output)

			w.Write([]byte(outputJson))
			return
		}

		for rows.Next() {
			rows.Scan(&bahasa.Id, &bahasa.Name)
			arr_bahasa = append(arr_bahasa, bahasa)
		}

		output.Status = 1
		output.Message = "Success"
		output.Data = arr_bahasa

		outputJson, _ := json.Marshal(output)
		w.Write([]byte(outputJson))

	} else {
		w.WriteHeader(500)
		w.Write([]byte("Wrong Method"))
	}
}
