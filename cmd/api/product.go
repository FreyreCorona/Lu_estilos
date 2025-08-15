package main

import (
	"errors"
	"net/http"
	"time"

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
	var input struct {
		Name         string                 `json:"name"`
		Description  *string                `json:"description,omitempty"`
		BarCode      *string                `json:"bar_code,omitempty"`
		Category     *string                `json:"category,omitempty"`
		InitialStock int32                  `json:"initial_stock"`
		ActualStock  int32                  `json:"actual_stock"`
		Price        float64                `json:"price"`
		DueDate      *time.Time             `json:"due_date,omitempty"`
		Images       []*models.ProductImage `json:"images,omitempty"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	product := &models.Product{
		Name:         input.Name,
		Description:  input.Description,
		BarCode:      input.BarCode,
		Category:     input.Category,
		InitialStock: input.InitialStock,
		ActualStock:  input.ActualStock,
		Price:        input.Price,
		DueDate:      input.DueDate,
		Images:       input.Images,
	}

	err = app.Models.Products.Insert(product)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	product, err = app.Models.Products.Get(product.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"product": product}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) putProduct(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
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
	var input struct {
		Name         *string
		Description  *string
		BarCode      *string
		Category     *string
		InitialStock *int32
		ActualStock  *int32
		Price        *float64
		DueDate      *time.Time
		Images       []*models.ProductImage
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// start validating if are nil for assignment
	if input.Name != nil {
		product.Name = *input.Name
	}
	if input.Description != nil {
		product.Description = input.Description
	}
	if input.BarCode != nil {
		product.BarCode = input.BarCode
	}
	if input.Category != nil {
		product.Category = input.Category
	}
	if input.InitialStock != nil {
		product.InitialStock = *input.InitialStock
	}
	if input.ActualStock != nil {
		product.ActualStock = *input.ActualStock
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.DueDate != nil {
		product.DueDate = input.DueDate
	}
	if input.Images != nil {
		product.Images = input.Images
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
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
