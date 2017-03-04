package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world , this is handlefunc")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/helloworldgo", helloWorld)

	server.ListenAndServe()
}
