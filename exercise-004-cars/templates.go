package main

import (
	"html/template"
	"net/http"
)

var (
	playT *template.Template
	joinT *template.Template
)

// SetupTemplates initializes two templates for rendering
func SetupTemplates() {
	jFilePath := RootDir() + "/templates/join.html"
	joinT = template.Must(template.ParseFiles(jFilePath))

	pFilePath := RootDir() + "/templates/play.html"
	playT = template.Must(template.ParseFiles(pFilePath))
}

// RenderPlay renders the play template with a username
func RenderPlay(w http.ResponseWriter, username string) {
	playT.Execute(w, ViewVehicles(username))
}

// RenderJoin renders the join template
func RenderJoin(w http.ResponseWriter) {
	joinT.Execute(w, nil)
}
