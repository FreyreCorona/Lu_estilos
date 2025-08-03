package main

import (
	"fmt"
	"net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
