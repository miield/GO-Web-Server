package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
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

func main() {

	// raw http.HandleFunc, this is how the server handles requests
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome here...")
	// })

	// serve the static file using the FileServer
	fileServer := http.FileServer(http.Dir("./static"))
	// handle the static file
	http.Handle("/", fileServer)

	// Handle the form path with http.HandleFunc
	http.HandleFunc("/form", formHandler)

	// Handle the hello path with http.HandleFunc
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// start the server & listen for incoming requests
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// error handling
		log.Fatal(err)
	}
}
