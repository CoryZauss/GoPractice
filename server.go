package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request ) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	} )

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

//! we deine two different handlers. these handlers respond to any HTTP requests that matchs the string pattern we define as the first parameter