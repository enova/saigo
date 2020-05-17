package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var urlStr string

func init() {
	setup("../")
	address := "localhost:8080"
	resource := "/home/"

	u, _ := url.ParseRequestURI(address)
	u.Path = resource
	urlStr = u.String()
}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", urlStr, nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	assert.Contains(w.Body.String(), "Signed up names")
}

func TestPost(t *testing.T) {
	assert := assert.New(t)

	data := url.Values{}
	data.Add("username", "ViPuL")

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	assert.Nil(err)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	home(w, req)
	home(w, req)
	home(w, req)

	assert.Contains(w.Body.String(), "vipul")
	assert.Contains(w.Body.String(), "3")
}
