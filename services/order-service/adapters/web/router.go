package web

import (
	"net/http"
)

func NewRouter(handler *WebHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", handler.HandleOrder)
	return mux
}
