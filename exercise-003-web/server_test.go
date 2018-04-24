package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	file = "test.json"
	setup()
}
func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	assert.Contains(w.Body.String(), "Name:")

	e := saveNames()
	assert.Nil(e)

	e = loadNames()
	assert.Nil(e)

	os.Remove(file)
	// there is plenty more that could be tested here, but I'm trying to get through saigo quickly
}
