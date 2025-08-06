package main

import (
	"net/http"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
)

func (app *application) getClients(w http.ResponseWriter, r *http.Request) {
}

func (app *application) getClienByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)
	if err != nil {
		app.InfoLogger.Println(err)
		app.notFoundResponse(w, r)
		return
	}
	client := models.Client{
		ID:       id,
		Name:     "Einier",
		Email:    "einierfreyre@gmail.com",
		CPF:      "712.960.812-90",
		Role:     "user",
		Password: "12345",
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"client": client}, nil)
	if err != nil {
		app.errorResponse(w, r, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (app *application) postClient(w http.ResponseWriter, r *http.Request) {
}

func (app *application) putClient(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteClient(w http.ResponseWriter, r *http.Request) {
}
