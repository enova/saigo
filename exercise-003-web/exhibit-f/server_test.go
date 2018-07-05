package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func init() {
	setup("../")
}

func TestHome(t *testing.T) {
	a := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	a.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	a.Contains(w.Body.String(), "Enter Name")
	a.Contains(w.Body.String(), "Sign-up Stats")
}

func TestSubmit(t *testing.T) {
	a := assert.New(t)
	v := url.Values{}
	v.Add("username", "randomperson")

	req, err := http.NewRequest("POST", "localhost:8080/home", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = v
	a.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	a.Contains(w.Body.String(), "randomperson")
}
