package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestLoginSuccessfull(t *testing.T) {
	ts := httptest.NewServer(setupRouter())
	defer ts.Close()

	values := url.Values{}
	values.Add("username", "user")
	values.Add("password", "password")

	resp, _ := http.PostForm(fmt.Sprintf("%s/login", ts.URL), values)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestLoginFaild(t *testing.T) {
	ts := httptest.NewServer(setupRouter())
	defer ts.Close()

	values := url.Values{}
	values.Add("username", "user")
	values.Add("password", "user")

	resp, _ := http.PostForm(fmt.Sprintf("%s/login", ts.URL), values)
	assert.Equal(t, 401, resp.StatusCode)
}
