package main

import (
	"encoding/json"
	"errors"
	"maps"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (app *application) readParamID(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id > 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	json = append(json, '\n')
	w.WriteHeader(status)
	maps.Copy(w.Header(), headers)
	w.Header().Set("Content-Type", "application-json")
	w.Write(json)

	return nil
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, message string, status int) {
	enve := envelope{"error": message}
	err := app.writeJSON(w, status, enve, nil)
	if err != nil {
		app.InfoLogger.Println(err)
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
