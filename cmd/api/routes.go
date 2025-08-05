package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	// general
	router.HandlerFunc(http.MethodGet, "/", app.homeHandle)
	// auth
	router.HandlerFunc(http.MethodGet, "/auth/register", app.authRegister)
	router.HandlerFunc(http.MethodGet, "/auth/login", app.authLogin)
	router.HandlerFunc(http.MethodGet, "/auth/refresh", app.authRefresh)
	// Clients
	router.HandlerFunc(http.MethodGet, "/client/", app.getClients)
	router.HandlerFunc(http.MethodGet, "/client/:id", app.getClienByID)
	router.HandlerFunc(http.MethodPost, "/client/:id", app.postClient)
	router.HandlerFunc(http.MethodPut, "/client/:id", app.putClient)
	router.HandlerFunc(http.MethodDelete, "/client/:id", app.deleteClient)
	// Orders
	router.HandlerFunc(http.MethodGet, "/order/", app.getClients)
	router.HandlerFunc(http.MethodGet, "/order/:id", app.getClienByID)
	router.HandlerFunc(http.MethodPost, "/order/:id", app.postClient)
	router.HandlerFunc(http.MethodPut, "/order/:id", app.putClient)
	router.HandlerFunc(http.MethodDelete, "/order/:id", app.deleteClient)
	return router
}

func (app *application) homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
