package main

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
