package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWord).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// getUsers returns all users
func helloWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello World!!"))
}
