package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Todo struct {
	Title *string `json:"title"`
}

var db *sql.DB

func main() {
	var err error
	db, err = newDB()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	defer db.Close()

	router := http.NewServeMux()
	router.HandleFunc("GET /api/todo", getTodos)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func newDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL must be set in the .env file")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(50)

	return db, err
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT title FROM todo`)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var title sql.NullString
		if err := rows.Scan(&title); err != nil {
			handleError(err, w, http.StatusInternalServerError)
			return
		}

		if title.Valid {
			todos = append(todos, Todo{Title: &title.String})
		} else {
			todos = append(todos, Todo{Title: nil})
		}
	}

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		handleError(err, w, http.StatusInternalServerError)
	}
}

func handleError(err error, w http.ResponseWriter, status int) {
	log.Printf("Error: %v", err)
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}
