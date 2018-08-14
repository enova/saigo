package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func init() {
	setup(".")
}

func TestHomeGet(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)
	w := httptest.NewRecorder()
	home(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(resp.StatusCode, 200)
	assert.Contains(string(body), "Count")
	assert.Contains(string(body), "Enter Name")
	assert.Contains(string(body), "Add Name")
}

func TestHomePost(t *testing.T) {
	assert := assert.New(t)

	data := url.Values{}
	data.Add("username", "Gandalf")

	req, err := http.NewRequest("POST", "localhost:8080/home", strings.NewReader(data.Encode()))
	assert.Nil(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	home(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(resp.StatusCode, 200)
	assert.Contains(string(body), "Gandalf")

}

func TestHomeDelete(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("DELETE", "localhost:8080/home", nil)
	assert.Nil(err)
	w := httptest.NewRecorder()
	home(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(resp.StatusCode, 200)
	assert.Contains(string(body), "invalid request method")
}
