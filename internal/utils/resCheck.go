package utils

import (
	"bytes"
	"log"
	"strconv"

	"github.com/knakk/sparql"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: test
SELECT ?s (EXISTS{?s <http://purl.org/dc/elements/1.1/title> ?o} AS ?b)
{
  ?s <http://purl.org/dc/elements/1.1/title> ?o .
  }
`

// SB is a struct to hold boolean check on a resource
type SB struct {
	S string
	B bool
}

// ResCheck see if a resource exist to we can offer an inbox for it
func ResCheck(r string) bool {
	repo, err := LDNDBConn()
	if err != nil {
		log.Printf("%s\n", err)
	}

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	q, err := bank.Prepare("test", r)
	if err != nil {
		log.Printf("%s\n", err)
	}

	res, err := repo.Query(q)
	if err != nil {
		log.Printf("%s\n", err)
	}

	bindings := res.Results.Bindings // map[string][]rdf.Term
	check, err := strconv.ParseBool(bindings[0]["b"].Value)
	if err != nil {
		return false
	}

	return check
}
