package app

import (
	"net/http"
)

type requestHandler struct {
	internalHandler http.Handler
}

func (h *requestHandler) applyCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding",
	)
}

func (h *requestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.applyCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	h.internalHandler.ServeHTTP(w, r)
}
