package main

import (
	"log"
	"net/http"
	bookroute "projects/api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/books", bookroute.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", bookroute.GetBook).Methods("GET")
	r.HandleFunc("/api/books", bookroute.CreateBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", bookroute.UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", bookroute.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3001", r))

}
