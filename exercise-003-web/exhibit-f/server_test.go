package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func init() {
	setup("../")
}

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	assert.Contains(w.Body.String(), "User Name")
}