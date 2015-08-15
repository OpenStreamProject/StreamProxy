package proxy

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

// Runs a GET request as the server for a URL
// This method only returns a byte array
func GetBytesForURL(url string) []byte {
	fmt.Println("Proxying: " + url)

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}


// Calls GetBytesFromURL and then converts the output
func GetStringForURL(url string) string {
	return string(GetBytesForURL(url)[:])
}