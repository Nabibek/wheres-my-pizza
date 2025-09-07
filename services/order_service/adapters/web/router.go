package web

import (
	"net/http"
)

func NewRouter(handler *WebHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", WebHandler.HandleOrder)
	return mux
}
