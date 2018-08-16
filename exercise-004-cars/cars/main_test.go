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

var testUsername = "Gandalf"

var testVehicle = Vehicle{Name: "Jeep", Speed: "fast"}
var testVehicles = Vehicles{}

func init() {
	setup("../templates")
}

func TestJoinGet(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/join", nil)
	assert.Nil(err)
	w := httptest.NewRecorder()
	join(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(resp.StatusCode, 200)
	assert.Contains(string(body), "goCars! Server")
	assert.Contains(string(body), "Enter Your Name")
	assert.Contains(string(body), "Join!")
}

func TestJoinPost(t *testing.T) {
	assert := assert.New(t)
	testUsername := "Gandalf"

	// set params to send as form values
	data := url.Values{}
	data.Add("username", testUsername)

	// send request with mocked form values
	req, err := http.NewRequest("POST", "localhost:8080/join", strings.NewReader(data.Encode()))
	assert.Nil(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	join(w, req)

	// validate cookie created correctly
	reqCookie := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}
	reqCookieUsername, err := reqCookie.Cookie("username")
	assert.Nil(err)
	assert.Equal(testUsername, reqCookieUsername.Value)

	// validate redirect code sent
	resp := w.Result()
	assert.Equal(resp.StatusCode, http.StatusSeeOther)
}

func TesJoinDelete(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("DELETE", "localhost:8080/join", nil)
	assert.Nil(err)
	w := httptest.NewRecorder()
	join(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(resp.StatusCode, 200)
	assert.Contains(string(body), "invalid request method")
}
