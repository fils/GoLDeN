package main

import (
	"log"
	"net/http"

	"../../internal/graph"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/id/ldn/{id}/inbox", graph.Headcall).Methods("HEAD")
	router.HandleFunc("/id/ldn/{id}/inbox", graph.Optionscall).Methods("OPTIONS")
	router.HandleFunc("/id/ldn/{id}/inbox", graph.Getcall).Methods("GET")
	router.HandleFunc("/id/ldn/{id}/inbox", graph.Postcall).Methods("POST")
	router.HandleFunc("/id/ldn/{id}/inbox/{nid}", graph.GetNotification).Methods("GET")

	log.Fatal(http.ListenAndServe(":6789", router))
}
