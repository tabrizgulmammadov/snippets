package main

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/tabriz-gulmammadov/snippets/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippets/{id}/view", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /users/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /users/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /users/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /users/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /snippets/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippets/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /users/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)
	return standard.Then(mux)
}
