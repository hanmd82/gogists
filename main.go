package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from GoGists"))
}

func showGist(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Display a specific gist..."))
}

func createGist(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Create a new gist..."))
}

func main() {
	// Start a new web server and register functions as handlers for respective URL patterns.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gists/create", createGist)

	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
