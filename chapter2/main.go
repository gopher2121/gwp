package main

import (
	"fmt"
	//	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//	files := http.FileServer(http.Dir("."))

	//	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func readThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is readThread Handler")
}

func postThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is postThread Handler")
}

func createThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is createThread Handler")
}

func newThread(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is newThread Handler")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is authenticate Handler")
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is signupAccount Handler")
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is signup  Handler")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is logout Handler")
}

func err(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is error Handler")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is login Handler")
}

func index(w http.ResponseWriter, r *http.Request) {

	//	templates := template.Must(template.ParseFiles(files...))

	//	templates.ExecuteTemplate(w, "index", "data")
	fmt.Fprintf(w, "hello world %s!", r.URL.Path[1:])
}
