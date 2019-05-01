package graph

import (
	"fmt"
	"log"
	"net/http"

	"../utils"
	"github.com/gorilla/mux"
)

// Getcall a POST call to the inbox
func Getcall(w http.ResponseWriter, r *http.Request) {
	log.Printf("Req: %s %s %s\n", r.URL.Scheme, r.Host, r.URL.Path)

	// Check if the resource exists
	if !utils.ResCheck(r.URL.Path) {
		log.Printf("Request of an unknown resource %s\n", r.URL.Path)
		// error out and return
		// return
	} else {
		log.Println("debug test to ensure true event here...")
	}

	log.Println("-----   Get Call no gid ----------")
	log.Println(r.Method)
	log.Println(r.URL.Path)
	log.Println(r.Header)
	//Content-Type: application/ld+json; charset=utf-8; profile="http://www.w3.org/ns/json-ld#expanded"
	w.Header().Add("Content-Type", "application/ld+json; charset=utf-8; profile=\"http://www.w3.org/ns/json-ld#expanded\"")
	jld := TestJSON()
	fmt.Fprintf(w, "%s", jld)
}

// GetNotification a POST call to the inbox
func GetNotification(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	fmt.Println(params["nid"])

	log.Printf("Req: %s %s %s\n", r.URL.Scheme, r.Host, r.URL.Path)

	// Check if the resource exists
	if !utils.ResCheck(r.URL.Path) {
		log.Printf("Request of an unknown resource %s\n", r.URL.Path)
		// error out and return
		// return
	} else {
		log.Println("debug test to ensure true event here...")
	}

	log.Println("-----   Get Call with gid ----------")
	log.Println(r.Method)
	log.Println(r.URL.Path)
	log.Println(r.Header)
	w.Header().Add("Content-Type", "application/ld+json")
	jld := TestJSON()
	fmt.Fprintf(w, "%s", jld)
}
