package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HelloHandler struct {
	log *log.Logger
}

func NewHelloHandler(l *log.Logger) *HelloHandler {
	return &HelloHandler{l}
}

func (h *HelloHandler) WriteBack(rw http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "something went wrong", http.StatusBadGateway)
		return
	}
	h.log.Printf("You have hit the request %s", r.URL.Path)
	fmt.Fprintf(rw, "You have sent, %s", data)
}

func (h *HelloHandler) DescribeFunc(rw http.ResponseWriter, r *http.Request) {
	h.log.Printf("You have hit the request %s", r.URL.Path)
	fmt.Fprintf(rw, "This is a handler function that greets you and returns the data you sent")
}
