package graph

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Headcall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	w.Header().Add("Allow", "GET, HEAD, OPTIONS, POST")
	w.Header().Add("Link", `<http://www.w3.org/ns/ldp#Resource>; rel="type", <http://www.w3.org/ns/ldp#RDFSource>; rel="type", <http://www.w3.org/ns/ldp#Container>; rel="type", <http://www.w3.org/ns/ldp#BasicContainer>; rel="type"`)
	w.Header().Add("Accept-Post", "application/ld+json, text/turtle")
	fmt.Fprintf(w, "")
}
