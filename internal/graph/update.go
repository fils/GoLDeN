package graph

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UpdateCall test out updates to Jena
func UpdateCall(insert string) ([]byte, error) {

	s := []byte(insert)

	host := "triplestore" // match name in compose file
	port := "7878"

	url := fmt.Sprintf("http://%s:%s/store", host, port) // TODO   add to a config file or CLI option.
	// fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(s))
	req.Header.Set("Content-Type", "application/n-quads")
	//	req.Header.Set("Content-Type", "application/sparql-update")

	fmt.Println(string(s))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println("UpdateCall response Status:", resp.Status)
	fmt.Println("UpdateCall response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("UpdateCall response Body:", string(body))

	return body, err
}
