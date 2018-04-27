package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestJoinGet(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080", nil)
	assert.Nil(err)

	recorder := httptest.NewRecorder()
	join(recorder, req)
	assert.Contains(recorder.Body.String(), "Enter Your Name")
}

func TestJoinPost(t *testing.T) {
	assert := assert.New(t)

	recorder := httptest.NewRecorder()

	testUsername := "DaveTest"
	form := url.Values{}
	form.Set("username", testUsername)

	req, err := http.NewRequest("POST", "localhost:8080", nil)
	req.Form = form

	assert.Nil(err)

	join(recorder, req)

	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}
	cookie, err := request.Cookie("username")
	assert.Nil(err)
	assert.Equal(testUsername, cookie.Value)
}
