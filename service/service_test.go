package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sharpvik/micro-go/database/names"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	e := New(names.NewRepo(nil)).Server()
	assert.Equal(t, http.StatusOK, response(t, get(t, "/ping"), e).StatusCode)
}

func response(t *testing.T, r *http.Request, h http.Handler) *http.Response {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w.Result()
}

func get(t *testing.T, url string) (req *http.Request) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err, "failed to create request:", url)
	return
}
