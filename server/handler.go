package server

import (
	"net/http"
)

type mainHandler struct {
}

func newMainHandler() (h *mainHandler) {
	return &mainHandler{}
}

func (h *mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Micro!"))
}
