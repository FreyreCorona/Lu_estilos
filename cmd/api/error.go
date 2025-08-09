package main

import "net/http"

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, message string, status int) {
	// wraps the response
	enve := envelope{"error": message}
	// Write as JSON
	err := app.writeJSON(w, status, enve, nil)
	if err != nil {
		app.Logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the server could not have that method"
	app.errorResponse(w, r, message, http.StatusNotFound)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := "the server could not support this method"
	app.errorResponse(w, r, message, http.StatusMethodNotAllowed)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, err.Error(), http.StatusBadRequest)
}
