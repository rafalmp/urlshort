package main

import (
	"fmt"
	"net/http"

	"github.com/rafalmp/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on port 8000")
	http.ListenAndServe(":8000", mapHandler)
}

func defaultMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", hello)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
