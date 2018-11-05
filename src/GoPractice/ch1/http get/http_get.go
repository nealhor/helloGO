package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ :=
		http.Get("http://example.com/") // makes an HTTP GET requst
	body, _ :=
		ioutil.ReadAll(resp.Body) // read te body from the response
	fmt.Println(string(body)) // prints body as a string
	resp.Body.Close()         // closes the connection

}
