package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/url"
	"strings"
)

func init() {
	setup("../")
}

func TestHomeGet(t *testing.T) {
	assert := assert.New(t)
	
	req, err := http.NewRequest("GET", "http://localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(resp.StatusCode, 200)
	assert.Contains(string(body), "User attempts")
}

func TestHomePost(t *testing.T) {
	assert := assert.New(t)
	data := url.Values{"username": []string{"dave"}}

	req, err := http.NewRequest("POST", "http://localhost:8080/home", strings.NewReader(data.Encode()))
	assert.Nil(err)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	home(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(resp.StatusCode, 200)
	assert.Contains(string(body), "dave")
	assert.Contains(string(body), "1")
}

func TestHomePut(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("PUT", "http://localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Contains(string(body), "Not a supported method")
}