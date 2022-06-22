package main

import (
	"log"
	"net/http"
)


func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}

//? " go run server.go"
//? http://localhost:8081/index.html

//!  How do we go about securing our web server and serving our content using HTTPS?

//* With Go, we can modify our existing web server to use http.ListenAndServeTLS

//? Generating Keys
/*
If you don’t have keys already generated, you can generate self-signed certs locally using openssl:

$ openssl genrsa -out server.key 2048
$ openssl ecparam -genkey -name secp384r1 -out server.key
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
This should generate self-signed certs for your locally and you can then try start your https web server by typing go run main.go. When you navigate to https://localhost:8081 you should see that the connection is now secured based on your self-signed cert.

*/