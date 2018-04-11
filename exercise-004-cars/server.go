package main

import (
	"fmt"
	"net/http"
)

func setup() {
	SetupDir()
	SetupTemplates()
	SetupHandlers()
}

func main() {
	setup()

	fs := http.FileServer(http.Dir(RootDir() + "/public/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	fmt.Println("Listening on :8086...")
	http.ListenAndServe(":8086", nil)
}
