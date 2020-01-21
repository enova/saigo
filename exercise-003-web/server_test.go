package main


import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init() {
	setup()
}

func TestServer(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)

	w := httptest.NewRecorder()
	home(w, req)
	assert.Contains(w.Body.String(), "Names Previously Submitted")
}

func TestSubmit(t *testing.T) {
	assert := assert.New(t)

	reader := strings.NewReader("name=Luke")
	req, _ := http.NewRequest("POST","/submit",reader)
	w := httptest.NewRecorder()
	submit(w,req)
	assert.Equal(w.Code,http.StatusSeeOther)

}

func TestSubmittingMultiple(t *testing.T) {
	assert := assert.New(t)
	reader := strings.NewReader("name=Luke")
	req, _ := http.NewRequest("POST","/submit",reader)
	w := httptest.NewRecorder()
	submit(w,req)
	w = httptest.NewRecorder()
	submit(w,req)

	hreq, err := http.NewRequest("GET", "localhost:8080/home", nil)
	assert.Nil(err)

	hw := httptest.NewRecorder()
	home(hw, hreq)
	assert.Contains(hw.Body.String(), "Luke</td><td>2")

}