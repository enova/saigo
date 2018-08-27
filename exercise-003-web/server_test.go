package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest(http.MethodGet, "localhost:8080/", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	index(w, req)
	assert.Contains(w.Body.String(), "Exercise")
}

func TestSubmitValid(t *testing.T) {
	assert := assert.New(t)

	form := url.Values{}
	form.Add("name", "carlos")
	payload := strings.NewReader(form.Encode())
	req, err := http.NewRequest(http.MethodPost, "localhost:8080/submit", payload)
	req.PostForm = form
	assert.Nil(err)

	w := httptest.NewRecorder()
	submit(w, req)
	assert.Equal(http.StatusSeeOther, w.Code)
	assert.Equal("/", w.HeaderMap.Get("location"))
}

func TestSubmitEmptyValue(t *testing.T) {
	assert := assert.New(t)

	form := url.Values{}
	form.Add("name", "")
	payload := strings.NewReader(form.Encode())
	req, err := http.NewRequest(http.MethodPost, "localhost:8080/submit", payload)
	req.PostForm = form
	assert.Nil(err)

	w := httptest.NewRecorder()
	submit(w, req)
	assert.Equal(http.StatusBadRequest, w.Code)
}

func TestSubmitEmptyPayload(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest(http.MethodPost, "localhost:8080/submit", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	submit(w, req)
	assert.Equal(http.StatusBadRequest, w.Code)
}
