package main

import (
	"database/sql"
	"net/http"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("tasks", "gabriel:password@tcp(127.0.0.1:3306)/tasks")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	http.HandleFunc("/api/tasks", CreateTask)
	http.ListenAndServe(":8080", nil)
}
