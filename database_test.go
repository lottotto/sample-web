package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestDataBaseIntegrate(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/list", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDatabasecannotConnect(t *testing.T) {
	os.Setenv("DbHost", "aaaaaaaa")

	var goenv Env
	envconfig.Process("", &goenv)
	fmt.Println(goenv.DbHost)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/list", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}
