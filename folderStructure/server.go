package main

import (
	"log"
	"net/http"
)


func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8081", nil))
}

//? " go run server.go"
//? http://localhost:8081/index.html

//! we’ve moved away from using the HandleFunc method and we’ve started using http.Handle() passing in our path and http.Dir() which points to our newly created static/ directory.