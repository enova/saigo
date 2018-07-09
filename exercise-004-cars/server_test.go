package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestJoinGet(t *testing.T) {

	req, _ := http.NewRequest("GET", "/join", nil)
	w := httptest.NewRecorder()
	join(w, req)

	bodyString := w.Body.String()
	assert.Contains(t, bodyString, "goCars!")
	assert.Contains(t, bodyString, "Copyright &copy; goCars! 2032. All rights thrown out the window.")
}

func TestJoinPost(t *testing.T) {
	username := "Aush"
	form := url.Values{}
	form.Add("username", username)

	req, _ := http.NewRequest("POST", "/join", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = form

	w := httptest.NewRecorder()
	join(w, req)

	assert.Equal(t, 302, w.Code)
	assert.Equal(t, "/play", w.HeaderMap["Location"][0])
	assert.Contains(t, w.HeaderMap["Set-Cookie"][0], "username=Aush")
}

func TestExit(t *testing.T) {
	req, _ := http.NewRequest("GET", "/exit", nil)
	w := httptest.NewRecorder()
	exit(w, req)

	assert.Equal(t, 302, w.Code)
	assert.Equal(t, "/join", w.HeaderMap["Location"][0])
	assert.Contains(t, w.HeaderMap["Set-Cookie"][0], "username=; Max-Age=0")
}
