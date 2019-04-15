package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	homeT = template.Must(template.ParseFiles("../exhibit-e/home.html"))
}

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	assert.Contains(w.Body.String(), "I would like to do")
}
