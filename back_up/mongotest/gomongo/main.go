package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	uri := os.Getenv("MONGOLAB_URL")
	if uri == "" {
		fmt.Println("no connection string provided ($MONGOLAB_URL)")
		os.Exit(1)
	}
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB("lyricaldb-mg")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/songs", getSongs).Methods("GET")
	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	songDAO := SongDAO{db}
	songs, err := songDAO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, songs)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
