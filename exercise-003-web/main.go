package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
)

func serve(){
	setup("./templates")
	restoreState(&users, DbPath)
	server := &http.Server{Addr: ":8080", Handler: buildHandlers()}


	go stop(server, func () chan os.Signal{
		stopSignal := make(chan os.Signal, 1)
		signal.Notify(stopSignal, os.Interrupt)
		return stopSignal
	}())
	server.ListenAndServe()
}

func stop(server *http.Server, stopSignal chan os.Signal){
	<-stopSignal
	saveState(&users, DbPath)
	server.Shutdown(context.Background())
}

func buildHandlers() * http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/users", handleUsers)
	return mux
}

func main() {
	serve()
}
