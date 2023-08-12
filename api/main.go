package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "gabriel:password@tcp(127.0.0.1:3306)/tasks")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	http.HandleFunc("/api/tasks", CreateTask)
	http.ListenAndServe(":8080", nil)
}
