package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 "gabriel",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "todoDB",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/tasks", tasksHandler)
	http.ListenAndServe(":8080", nil)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rows, err := db.Query("SELECT id, title, description, done FROM tasks")
		if err != nil {
			http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tasks []Task
		for rows.Next() {
			var task Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
				http.Error(w, "Failed to scan task", http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	case "POST":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res, err := db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
		if err != nil {
			http.Error(w, "Failed to insert task", http.StatusInternalServerError)
			return
		}

		id, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve last inserted ID", http.StatusInternalServerError)
			return
		}
		task.ID = int(id)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	case "PUT":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if _, err := db.Exec("UPDATE tasks SET title=?, description=?, done=? WHERE id=?", task.Title, task.Description, task.Done, task.ID); err != nil {
			http.Error(w, "Failed to update task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case "DELETE":
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if _, err := db.Exec("DELETE FROM tasks WHERE id=?", task.ID); err != nil {
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
