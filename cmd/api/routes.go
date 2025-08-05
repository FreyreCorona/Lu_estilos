package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.homeHandle)
	router.HandlerFunc(http.MethodGet, "/auth/register", app.authRegister)
	router.HandlerFunc(http.MethodGet, "/auth/login", app.authLogin)
	router.HandlerFunc(http.MethodGet, "/auth/refresh", app.authRefresh)

	return router
}

func (app *application) homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
