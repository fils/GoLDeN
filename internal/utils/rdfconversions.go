package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/piprate/json-gold/ld"
)

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

func JSONLDToNQ(jsonld string) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1
	options.Format = "application/n-quads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
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
