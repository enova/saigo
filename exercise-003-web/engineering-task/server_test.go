package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	index(w, req)
	assert.Contains(w.Body.String(), "Sign-Up")
}

func TestSignup(t *testing.T) {
	assert := assert.New(t)

	form := url.Values{}
	form.Set("username", "asdf")
	body := bytes.NewBufferString(form.Encode())

	req, err := http.NewRequest("POST", "localhost:8080/signup", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(err)

	w := httptest.NewRecorder()
	addUser(w, req)
	assert.Contains(w.Body.String(), "asdf")
}
