package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	// error
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// general
	router.HandlerFunc(http.MethodGet, "/", app.homeHandle)
	// auth
	router.HandlerFunc(http.MethodGet, "/auth/register", app.authRegister)
	router.HandlerFunc(http.MethodGet, "/auth/login", app.authLogin)
	router.HandlerFunc(http.MethodGet, "/auth/refresh", app.authRefresh)
	router.HandlerFunc(http.MethodPost, "/auth/logout", app.authLogout)
	// Clients
	router.HandlerFunc(http.MethodGet, "/client/:id", app.getClienByID)
	router.HandlerFunc(http.MethodPost, "/client", app.postClient)
	router.HandlerFunc(http.MethodPatch, "/client/:id", app.putClient)
	router.HandlerFunc(http.MethodDelete, "/client/:id", app.deleteClient)
	// Orders
	router.HandlerFunc(http.MethodGet, "/order/", app.getOrders)
	router.HandlerFunc(http.MethodGet, "/order/:id", app.getOrderByID)
	router.HandlerFunc(http.MethodPost, "/order/:id", app.postOrder)
	router.HandlerFunc(http.MethodPut, "/order/:id", app.putOrder)
	router.HandlerFunc(http.MethodDelete, "/order/:id", app.deleteOrder)
	// Product
	router.HandlerFunc(http.MethodGet, "/product/:id", app.getProductByID)
	router.HandlerFunc(http.MethodPost, "/product/:id", app.postProduct)
	router.HandlerFunc(http.MethodPut, "/product/:id", app.putProduct)
	router.HandlerFunc(http.MethodDelete, "/product/:id", app.deleteProduct)

	return router
}

func (app *application) homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
