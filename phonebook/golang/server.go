package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"html"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

var collection *mgo.Collection

func init() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	collection = session.DB("test").C("person")
}

type Person struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstname" bson:"firstName"`
	LastName  string        `json:"lastname" bson:"lastName"`
	Phone     string        `json:"phone" bson:"phone"`
	Email     string        `json:"email" bson":"email"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	var results []Person
	err := collection.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(results)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/people", PersonHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
