package graph

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../utils"
	"github.com/gorilla/mux"
)

func Postcall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on body parameter read %v \n", err)
		log.Println(err)
		//	response.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "")
		return
	}

	// Check that it is JSON-LD and attempt to convert it to Turtle
	nq, err := utils.JSONLDToNQ(string(body))
	if err != nil {
		log.Println(err)
	}

	// TODO
	// The ID of this URL needs to be resolved
	// use a sha hash?  It would be a big URL
	// would work though....

	insert := "INSERT DATA {" + nq + "}"

	// Try to insert into Jena
	_, err = UpdateCall([]byte(insert))
	if err != nil {
		log.Printf("Error on updatecall %v \n", err)
		log.Println(err)
		//response.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "")
		return
	}

	w.Header().Add("Location", "http://c254046f.ngrok.io/id/ldn/1234/inbox/1")
	w.WriteHeader(201)
	fmt.Fprintf(w, "")

}
