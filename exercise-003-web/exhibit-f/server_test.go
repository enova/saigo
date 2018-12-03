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

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "http://localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	assert.Contains(w.Body.String(), "Enter Name")
	assert.Contains(w.Body.String(), "Username")
	assert.Contains(w.Body.String(), "Attempts")
}

func TestPost(t *testing.T) {
	assert := assert.New(t)

	v := url.Values{}

	req, err := http.NewRequest("POST", "http://localhost:8080/home", nil)
	v.Add("username", "Kavya")
	req.PostForm = v
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	assert.Contains(w.Body.String(), "Kavya")
	assert.Contains(w.Body.String(), "1")
}

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("PUT", "http://localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	assert.Contains(w.Body.String(), "Invalid HTTP method")

}
