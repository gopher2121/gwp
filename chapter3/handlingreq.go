package main

import (
	"fmt"
	"net/http"
)

//MyHandler is an empty struct that has a method ServeHTTP , so it satisfies the handler interface
type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pageName := "myHandler"
	fmt.Fprintf(w, "Hello , welcome to  %s page", pageName)
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
