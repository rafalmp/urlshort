package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the server on port 8000")
	http.ListenAndServe(":8000", defaultMux())
}

func defaultMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", hello)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
