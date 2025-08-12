package main

import (
	"errors"
	"net/http"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
)

func (app *application) getClienByID(w http.ResponseWriter, r *http.Request) {
	// get the id from the URL
	id, err := app.readParamID(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// build the client data
	client, err := app.Models.Clients.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// show as JSON format
	err = app.writeJSON(w, http.StatusOK, envelope{"client": client}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) postClient(w http.ResponseWriter, r *http.Request) {
	// Define a input struct with the expected data from the request
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		CPF      string `json:"cpf"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	// try to decode te request on the addres of the input struct
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// hash the password
	hashed, err := models.HashPassword(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	input.Password = hashed
	// map the input fields to an client object
	client := &models.Client{
		Name:     input.Name,
		Email:    input.Email,
		CPF:      input.CPF,
		Password: input.Password,
		Role:     input.Role,
	}
	// insert on the database
	err = app.Models.Clients.Insert(client)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	client, err = app.Models.Clients.Get(client.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// write the writed client in a JSON format and return the 201 status code
	err = app.writeJSON(w, http.StatusCreated, envelope{"client": client}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) putClient(w http.ResponseWriter, r *http.Request) {
	// get the id from parameters
	id, err := app.readParamID(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// retrieve the actual user from the database
	client, err := app.Models.Clients.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// create a input struct for get the request (all fields are pointers for easy check if are nil values)
	var input struct {
		Name     *string
		Email    *string
		CPF      *string
		password *string
		Role     *string
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	if input.Name != nil {
		client.Name = *input.Name
	}
	if input.Email != nil {
		client.Email = *input.Email
	}
	if input.CPF != nil {
		client.CPF = *input.CPF
	}
	if input.password != nil {
		hashed, err := models.HashPassword(*input.password)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
		client.Password = hashed
	}
	if input.Role != nil {
		client.Role = *input.Role
	}
	err = app.Models.Clients.Update(client)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"client": client}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteClient(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.Models.Clients.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "client successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
