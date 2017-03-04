package main

import (
	"fmt"
	"net/http"
)

type SampleHandler struct{}

func (s *SampleHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "sample handler")
}

func main() {

	mux := http.NewServeMux()
	sh := SampleHandler{}
	mux.Handle("/", &sh)

	//http.ListenAndServe(":8080",nil)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()

}
