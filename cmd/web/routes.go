package main

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler { //dont forget that astric refers to the type or implmentation

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	}) // this will make router return the same method even if handler func not exists

	// fileServer := http.FileServer(http.Dir("./ui/static"))

	fileServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the correct MIME type for CSS and JS files
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
		} else if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		}
		// Serve the file
		http.FileServer(http.Dir("./ui/static")).ServeHTTP(w, r)
	})

	router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	middlewares := alice.New(app.recoverPanic, app.logRequest, app.secureHeaders)
	return middlewares.Then(router)
	// very clean yo use  justinas/alice package also has append function

	// important to mention ' conflict router isnt allowed , but gorilla/mux allow it
}
