package main

import (
	"fmt"

	"../../internal/consumer"
)

func main() {
	targetA := "https://linkedresearch.org/ldn/tests/discover-inbox-link-header"
	targetB := "https://linkedresearch.org/ldn/tests/discover-inbox-rdf-body"
	fmt.Printf("LDN test URLs: \n targetA: %s \n targetB: %s \n", targetA, targetB)

	consumer.ReadGetBody(targetA)
}
