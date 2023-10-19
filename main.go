package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// raw http.HandleFunc
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome here...")
	// })

	// Handle the http.HandleFunc
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// start the server & listen for incoming requests
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// error handling
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// checks the req path
	if r.URL.Path != "/hello" {
		// set resq w to 404
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// checks the req method, GET
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Welcome here...")
}