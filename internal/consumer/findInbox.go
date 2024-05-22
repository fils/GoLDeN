package consumer

import (
	"fmt"
	"net/http"
)

// ReadGetBody read the get body..   duh...
func ReadGetBody(url string) string {
	r, _ := http.Get(url)
	m := r.Header

	for k := range m {
		fmt.Printf("key[%s] value[%s]\n", k, m[k])
	}
	// Look for link then make a map of the comma separated ; based map items

	err := r.Body.Close()
	if err != nil {
		return "Error Closing Body"
	}
	return "URL of inbox"
}
