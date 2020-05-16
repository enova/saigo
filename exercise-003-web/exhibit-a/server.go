package main

import (
	"fmt"
	"net/http"
)

// router method for '/hello'
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just wanted to say hi!")
}

func main() {
	http.HandleFunc("/", hello)       // register a new router with a matcher for the URL path
	http.ListenAndServe(":8080", nil) // serve a webserver listening on port 8080
}
