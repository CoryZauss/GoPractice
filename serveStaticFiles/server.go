package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w func(http.ResponseWriter, *http.Request){
		http.ServeFile(w, r, r.URL.Path[1:])
	}))

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

//? " go run server.go"
//? http://localhost:8081/index.html