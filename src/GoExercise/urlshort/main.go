package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml":     "https://godoc.org/gopkg.in/yaml.v2",
		"/y":        "https://www.youtube.com/",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, list of url")
	fmt.Fprintln(w, "/y - youtube.com")
	fmt.Fprintln(w, "/urlshort - https://godoc.org/github.com/gophercises/urlshort")
	fmt.Fprintln(w, "/yaml - https://godoc.org/gopkg.in/yaml.v2")
	fmt.Fprintln(w, "/urlshort-final - https://github.com/gophercises/urlshort/tree/solution")

	//fmt.Fprintln(w, "Hello, world!")
}
