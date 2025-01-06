package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static"))

	//router.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileServer))   //  errors
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave, app.noSurf)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)
	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))
	router.Handler(http.MethodPost, "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Handler(http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	// adding more staff to protect login access
	router.Handler(http.MethodGet, "/user/login", protected.ThenFunc(app.userLogin))

	standard := alice.New(app.recoverPanic, app.logRequest, app.secureHeaders)
	return standard.Then(router)

}
