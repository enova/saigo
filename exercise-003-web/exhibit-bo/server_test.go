package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	setup("../")
}

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	index(w, req)
	assert.Contains(w.Body.String(), "Names")
}

func TestNames(t *testing.T) {
	assert := assert.New(t)
	form := url.Values{}
	form.Add("firstName", "Bob")
	form.Add("lastName", "Loblaw")
	req, err := http.NewRequest("POST", "localhost:8080/names", strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	assert.Nil(err)

	w := httptest.NewRecorder()
	names(w, req)
	assert.Contains(w.Body.String(), "Bob Loblaw")
}
