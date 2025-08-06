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
	// get the param from the context of the request
	params := httprouter.ParamsFromContext(r.Context())
	// try to convert to int64
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// convert object to json
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// append new line
	json = append(json, '\n')
	// set custom headers if exist
	maps.Copy(w.Header(), headers)
	// set the content type for display
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)

	return nil
}
