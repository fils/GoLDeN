package utils

import (
	"log"
	"time"

	"github.com/knakk/sparql"
)

func LDNDBConn() (*sparql.Repo, error) {
	repo, err := sparql.NewRepo("http://localhost:3030/ldn/sparql",
		sparql.Timeout(time.Millisecond*15000),
	)
	if err != nil {
		log.Printf("%s\n", err)
	}
	return repo, err
}
