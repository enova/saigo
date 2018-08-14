package main

import (
	"fmt"
	"net/http" // provides HTTP client and server implementations
)

func hello(w http.ResponseWriter, r *http.Request) { // w = Interface used by HTTP handler to construct HTTP response
	fmt.Fprintf(w, "Just wanted to say hi!") // r = HTTP request received by a server or to be sent by a client
}

func main() {
	http.HandleFunc("/", hello)       // routes "/" traffic to hello func, passing (ResponseWriter, *Request)
	http.ListenAndServe(":8080", nil) // listens on port 8080, calls Serve with handler (nil, typically) to handle requests
}
