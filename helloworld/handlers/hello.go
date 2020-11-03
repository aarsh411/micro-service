package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	l.h.Println("hello world")
	d, err := ioutil.ReadAll(r.Body)
	fmt.Fprintf(rw, "Hello %s", d)
	if err != nil {
		http.Error(rw, "error", http.StatusBadRequest)
		return
	}
	fmt.Printf(rw, "Hello %s", d)
}
