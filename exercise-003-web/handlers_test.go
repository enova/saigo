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
	setup("./templates")
}

func buildPostUsersRequest(data *url.Values) (* httptest.ResponseRecorder, *http.Request){
	req, _ := http.NewRequest("POST", "localhost:8080/users", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return httptest.NewRecorder(), req
}

func TestHandleRoot(t *testing.T) {
	assert := assert.New(t)
	req, _ := http.NewRequest("GET", "localhost:8080/home", nil)
	w := httptest.NewRecorder()
	handleRoot(w, req)
	const successHttpCode = 200
	const successString = "Learning Go"
	assert.Equal(w.Code, successHttpCode)
	assert.Contains(w.Body.String(), successString)
}

func TestHandlePostUsersBlank(t *testing.T) {
	assert := assert.New(t)
	data := url.Values{"username": {""}}
	assert.Equal(len(users), 0)
	w, req := buildPostUsersRequest(&data)
	handleUsers(w, req)

	assert.Equal(w.Code, 400)
	assert.Equal(len(users), 0)
}

func TestHandlePostUsersValid(t *testing.T) {
	assert := assert.New(t)
	data := url.Values{"username": {"test"}}
	w, req := buildPostUsersRequest(&data)
	handleUsers(w, req)
	handleUsers(w, req)
	assert.Equal(w.Code, 307)
	assert.Equal(users["test"], 2)
}
