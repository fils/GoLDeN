package graph

import (
	"bytes"
	"log"

	"github.com/fils/GoLDeN/internal/utils"
	sparql "github.com/knakk/sparql"
)

const inboxqueries = `
# Comments are ignored, except those tagging a query.

# tag: test
prefix schema: <http://schema.org/>
prefix bds: <http://www.bigdata.com/rdf/search#>
select *
where {
  { ?s ?p ?o }
  UNION { GRAPH ?g { ?s ?p ?o } }
}
LIMIT 10
`

// SPO is a basic query
type SPO struct {
	S string
	P string
	O string
}

// URLSet is used for what?
type URLSet []string

// InboxGet used during development work
func InboxGet(r string) []SPO {
	repo, err := utils.LDNDBConn()
	if err != nil {
		log.Printf("%s\n", err)
	}

	f := bytes.NewBufferString(inboxqueries)
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
	rra := []SPO{}

	for _, i := range bindings {
		rr := SPO{S: i["s"].Value,
			P: i["p"].Value,
			O: i["o"].Value}
		rra = append(rra, rr)
	}

	return rra
}

// InboxGetRes get a resource from the inbox
func InboxGetRes(r string) []SPO {
	repo, err := utils.LDNDBConn()
	if err != nil {
		log.Printf("%s\n", err)
	}

	f := bytes.NewBufferString(inboxqueries)
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
	rra := []SPO{}

	for _, i := range bindings {
		rr := SPO{S: i["s"].Value,
			P: i["p"].Value,
			O: i["o"].Value}
		rra = append(rra, rr)
	}

	return rra
}

// TestJSON is an example JSON response package
func TestJSON() string {

	jld := `[
		{
		  "@id": "http://www.w3.org/ns/ldp#BasicContainer",
		  "@type": [
			"http://www.w3.org/2002/07/owl#Class"
		  ]
		},
		{
		  "@id": "http://www.w3.org/ns/ldp#RDFSource",
		  "@type": [
			"http://www.w3.org/2002/07/owl#Class"
		  ]
		},
		{
		  "@id": "http://cor.esipfed.org/ont/ldn/inbox",
		  "@type": [
			"http://www.w3.org/2002/07/owl#NamedIndividual",
			"http://www.w3.org/2002/07/owl#Ontology",
			"http://www.w3.org/ns/ldp#Container",
			"http://www.w3.org/ns/ldp#Resource",
			"http://www.w3.org/ns/ldp#RDFSource",
			"http://www.w3.org/ns/ldp#BasicContainer"
		  ],
		  "http://www.w3.org/2000/01/rdf-schema#comment": [
			{
			  "@type": "http://www.w3.org/2000/01/rdf-schema#Literal",
			  "@value": "The ESIP Linked Data Notifications (LDN) Inbox graph represents a registry of LDN submitted to and received through the ESIP LDN service at http://cor.esipfed.org/ldn/inbox/.\nThe ESIP LDN service uses COR as a storage backend to persist and serve LDN for the ESIP community."
			}
		  ],
		  "http://www.w3.org/2000/01/rdf-schema#label": [
			{
			  "@type": "http://www.w3.org/2000/01/rdf-schema#Literal",
			  "@value": "ESIP Linked Data Notifications Inbox Graph"
			}
		  ],
		  "http://www.w3.org/2000/01/rdf-schema#seeAlso": [
			{
			  "@type": "http://www.w3.org/2000/01/rdf-schema#Literal",
			  "@value": "https://linkedresearch.org/ldn/"
			},
			{
			  "@type": "http://www.w3.org/2000/01/rdf-schema#Literal",
			  "@value": "https://github.com/esipfed/pyldn"
			}
		  ]
		},
		{
		  "@id": "http://www.w3.org/ns/ldp#Resource",
		  "@type": [
			"http://www.w3.org/2002/07/owl#Class"
		  ]
		},
		{
		  "@id": "http://www.w3.org/ns/ldp#Container",
		  "@type": [
			"http://www.w3.org/2002/07/owl#Class"
		  ]
		},
		{
		  "@id": "http://cor.esipfed.org/ldn/inbox/",
		  "@type": [
			"http://www.w3.org/ns/ldp#BasicContainer",
			"http://www.w3.org/ns/ldp#Resource",
			"http://www.w3.org/ns/ldp#Container",
			"http://www.w3.org/ns/ldp#RDFSource"
		  ],
		  "http://www.w3.org/ns/ldp#contains": [
			{
			  "@id": "http://pala.serveo.net/api/v1/graph/1234/inbox/1"
			}
		  ]
		}
	  ]

	`
	return jld
}
