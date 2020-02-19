package main

import (
	"database/sql"
	"fmt"
	"latihan/routing/controller"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := connectDb()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controller.IndexHandler(w, r)
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		controller.PostHandler(w, r)
	})

	http.HandleFunc("/post_json", func(w http.ResponseWriter, r *http.Request) {
		controller.PostJsonHandler(w, r)
	})

	http.HandleFunc("/get_data", func(w http.ResponseWriter, r *http.Request) {
		controller.GetDataHandler(db, w, r)
	})

	fmt.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	fmt.Println(err.Error())
}

func connectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/latihan")
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}
