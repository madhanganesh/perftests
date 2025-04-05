package main

import (
	"database/sql"
	"log"
	"net/http"
	"todopad/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	templ, err := utils.LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	var router = http.NewServeMux()
	router.HandleFunc("GET /", index(templ, db))
	router.HandleFunc("GET /about", about(templ))

	log.Fatal(http.ListenAndServe(":8000", router))
}

type Todo struct {
	ID    int
	Title string
}

func index(templ *utils.TemplCache, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := getTodos(db)
		if err != nil {
			handleError(w, "error getting todos", err)
		}
		err = templ.Render(w, "index.html", todos)
		if err != nil {
			handleError(w, "error rendering", err)
		}
	}
}

func about(templ *utils.TemplCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templ.Render(w, "about.html", nil)
		if err != nil {
			handleError(w, "error rendering", err)
		}
	}
}

func handleError(w http.ResponseWriter, message string, err error) {
	log.Printf("%s: %v", message, err)
	http.Error(w, message, http.StatusInternalServerError)
}

func getTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, title FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil

}
