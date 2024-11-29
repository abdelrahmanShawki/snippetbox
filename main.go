package main

import (
	"log"
	"net/http"
)

func home(writer http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(writer, r)
		return
	}
	writer.Write([]byte("Guten Tag !"))
}

func snippetView(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("snippt view"))
}

func snippetCreate(writer http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {

		writer.Header().Set("Allow", "Post")
		writer.WriteHeader(405)
		writer.Write([]byte("Method not allowed"))
		return
	}
	writer.Write([]byte("snippt Creat "))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("server is running on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
