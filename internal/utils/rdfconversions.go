package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/piprate/json-gold/ld"
)

// ContextMapping holds the JSON-LD mappings for cached context
type ContextMapping struct {
	Prefix string
	File   string
}

func JSONLDToNQ(jsonld string) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	// use local assets for context files to avoid network calls
	client := &http.Client{}
	nl := ld.NewDefaultDocumentLoader(client)

	m := make(map[string]string)
	m["https://schema.org/"] = "./assets/schemaorg-current-https.jsonld"
	m["http://schema.org/"] = "./assets/schemaorg-current-http.jsonld"

	cdl := ld.NewCachingDocumentLoader(nl)
	err := cdl.PreloadWithMapping(m)
	if err != nil {
		log.Println(err)
	}
	options.DocumentLoader = cdl

	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"

	var myInterface interface{}
	err = json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
		return "", err
	}

	triples, err := proc.ToRDF(myInterface, options) // returns triples but toss them, just validating
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return "", err
	}

	return fmt.Sprintf("%v", triples), err
}

func NQToJSONLD(triples string) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1

	doc, err := proc.FromRDF(triples, options)
	if err != nil {
		panic(err)
	}

	// ld.PrintDocument("JSON-LD output", doc)
	b, err := json.MarshalIndent(doc, "", " ")

	return b, err
}
