package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitServerHandlers(t *testing.T) {
	assert := assert.New(t)

	r1, err1 := http.NewRequest("GET", "localhost:8080/", nil)
	r2, err2 := http.NewRequest("GET", "localhost:8080/signup", nil)
	assert.Nil(err2)
	assert.Nil(err1)

	w := httptest.NewRecorder()
	signup(w, r1)
	assert.Contains(w.Body.String(), "Catalog")
	signup(w, r2)
	assert.Contains(w.Body.String(), "Catalog")
}
