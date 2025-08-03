package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
)

type Application struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	Clients     *models.ClietModel
}

func routes(m *http.ServeMux) {
	m.HandleFunc("GET /", homeHandle)
	m.HandleFunc("POST /auth/register", authRegister)
	m.HandleFunc("POST /auth/login", authLogin)
	m.HandleFunc("POST /auth/refresh", authRefresh)
}

func main() {
	Infolog := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate)
	Errorlog := log.New(os.Stdout, "ERROR: ", log.Ltime|log.Ldate|log.Lshortfile)

	var app Application
	app.ErrorLogger = Errorlog
	app.InfoLogger = Infolog

	mux := http.NewServeMux()
	routes(mux)

	server := http.Server{
		Addr:     ":8000",
		Handler:  mux,
		ErrorLog: app.ErrorLogger,
	}

	app.InfoLogger.Println("Starting Server...")
	Errorlog.Fatal(server.ListenAndServe())
}
