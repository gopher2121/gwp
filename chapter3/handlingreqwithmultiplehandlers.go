package main

import (
	"fmt"
	"net/http"
)

//HelloHandler has a ServeHTTP method
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Handler")
}

//WorldHandler has a ServeHTTP method
type WorldHandler struct{}

func (wh *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world handler")
}

func main() {
	hh := HelloHandler{}
	wh := WorldHandler{}

	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/hello", &hh)
	http.Handle("/world", &wh)

	server.ListenAndServe()
}
