package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	// "saigo/exercise-003-web/exhibit-e"
)

func init() {
	setup("../")
}

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/home", nil) // returns incoming Request, suitable for passing to an http.Handler for testing.
	assert.Nil(err)

	w := httptest.NewRecorder() // returns an initialized ResponseRecorder
	home(w, req)                // render home.html
	assert.Contains(w.Body.String(), "I would like to do")
}
