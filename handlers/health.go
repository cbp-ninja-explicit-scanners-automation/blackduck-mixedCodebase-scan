package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type HealthHandler struct {
	log *log.Logger
}

func NewHealthHandler(l *log.Logger) *HealthHandler {
	return &HealthHandler{l}
}

func (h *HealthHandler) HealthCheck(rw http.ResponseWriter, r *http.Request) {
	h.log.Printf("You have hit the request %s", r.URL.Path)
	fmt.Fprintf(rw, "Health check is working")
}
