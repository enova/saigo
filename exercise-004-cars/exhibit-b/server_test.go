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
	setup("../")
}

func TestJoinGet(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "http://localhost:8080/join", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	join(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Contains(string(body), "Join Server")
}

func TestJoinPost(t *testing.T) {
	assert := assert.New(t)
	username := "dave"
	data := url.Values{"username": []string{username}}

	req, err := http.NewRequest("POST", "http://localhost:8080/join", strings.NewReader(data.Encode()))
	assert.Nil(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	join(w, req)

	cookieReq := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
	cookie, err := cookieReq.Cookie("username")
	assert.Nil(err)
	assert.Equal(cookie.Value, username)
}

func TestPlayWithoutCookie(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "http://localhost:8080/play", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	play(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Contains(string(body), "Could not find a cookie named 'username'")
}

func TestPlayWithCookie(t *testing.T) {
	assert := assert.New(t)
	username := "dave"

	req, err := http.NewRequest("GET", "http://localhost:8080/play", nil)
	cookie := http.Cookie{Name: "username", Value: username}
	req.AddCookie(&cookie)
	assert.Nil(err)

	w := httptest.NewRecorder()
	play(w, req)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Contains(string(body), "Connected as")
}
