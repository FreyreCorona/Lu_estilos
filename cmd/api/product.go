package main

import (
	"errors"
	"net/http"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
)

func (app *application) getProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	product, err := app.Models.Products.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) postProduct(w http.ResponseWriter, r *http.Request) {
}

func (app *application) putProduct(w http.ResponseWriter, r *http.Request) {
}

func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.Models.Products.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Product successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
