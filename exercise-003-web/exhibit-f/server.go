package main

import (
	"context"
	"fmt"
	"net/http"
)

var server http.Server

func stopServer() {
	fmt.Println("Attempting server shutdown...")
	server.Shutdown(context.Background())
}

func startServer() {
	InitServerHandlers()

	fmt.Println("Listening on port 8080...")
	server = http.Server{Addr: ":8080", Handler: nil}
	server.ListenAndServe()
}

func main() {
	InitInterruptHandler()
	
	LoadSession()
	go startServer()
	<-ServerInterrupt
	go stopServer()
	SaveSession()
}
