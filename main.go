package main

import (
	"net/http"
	"log"
)

// Define a home handler function which writes a byte slice containing
// "Hello from GiGists" as the response body.

func home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello from GoGists"))
}

func main() {
	// Start a new web server
	// Register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}