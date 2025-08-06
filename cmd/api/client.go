package main

import (
	"fmt"
	"net/http"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
)

func (app *application) getClients(w http.ResponseWriter, r *http.Request) {
}

func (app *application) getClienByID(w http.ResponseWriter, r *http.Request) {
	// get the id from the URL
	id, err := app.readParamID(r)
	if err != nil {
		app.InfoLogger.Println(err)
		app.notFoundResponse(w, r)
		return
	}
	// build the client data
	client := models.Client{
		ID:       id,
		Name:     "Einier",
		Email:    "einierfreyre@gmail.com",
		CPF:      "712.960.812-90",
		Role:     "user",
		Password: "12345",
	}
	// show as JSON format
	err = app.writeJSON(w, http.StatusOK, envelope{"client": client}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
}

func (app *application) postClient(w http.ResponseWriter, r *http.Request) {
	// Define a input struct with the expected data from the request
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		CPF      string `json:"cpf"`
		Role     string `json:"role"`
		Password string `json:"password"`
	}
	// try to decode te request on the addres of the input struct
	err := app.readJSON(w, r, input)
	if err != nil {
		app.errorResponse(w, r, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// show on screen
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) putClient(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteClient(w http.ResponseWriter, r *http.Request) {
}
