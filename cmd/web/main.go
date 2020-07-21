package main

import (
	"log"
	"net/http"
)

func main() {
	// Start a new web server and register functions as handlers for respective URL patterns.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/gist", showGist)
	mux.HandleFunc("/gists/create", createGist)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
