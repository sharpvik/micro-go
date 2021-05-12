package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	handler := newMainHandler()
	assert.Equal(t, http.StatusOK, Response(t, Get(t, "/"), handler).StatusCode)
}

func Response(t *testing.T, r *http.Request, h http.Handler) *http.Response {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w.Result()
}

func Get(t *testing.T, url string) (req *http.Request) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	assert.NoError(t, err, "failed to create request:", url)
	return
}
