package graph

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/fils/GoLDeN/internal/utils"
	"github.com/gorilla/mux"
)

// Postcall posts LDN notificatons to a triplestore
func Postcall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on body parameter read %v \n", err)
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unprocessable entity"))
		return
	}

	// Check that it is JSON-LD and attempt to convert it to Turtle
	nq, err := utils.JSONLDToNQ(string(body))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Unprocessable entity"))
		return
	}

	// TODO
	// Make an XID
	// make a triple and add to the nq with the notification ID URI
	// Could also store to object store

	//	insert := "INSERT DATA {" + nq + "}"
	insert := nq // in this case we are using sparql over http not a sparql update call

	// Try to insert into Jena
	// TODO, update this to use Oxigraph
	_, err = UpdateCall(insert)
	if err != nil {
		log.Printf("Error on update call %v \n", err)
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unprocessable entity"))
		return
		//w.Header().Add("Location", "http://openknowledge.network/id/ldn/1234/inbox/1")
		return
	} else {
		// TODO get named graph value for graph and return that in the URL
		h := hash(insert)
		hs := strconv.FormatUint(uint64(h), 10)
		w.Header().Add("Location", fmt.Sprintf("http://mm.oceaninfohub.org/id/ldn/NAMEDGRAPH/inbox/%s", hs))
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "")
	}

}
