package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
)

type application struct {
	InfoLogger *log.Logger
	Clients    *models.ClietModel
}

func main() {
	Infolog := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate)

	app := application{
		InfoLogger: Infolog,
	}

	server := &http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	Infolog.Println("Starting Server...")
	log.Fatal(server.ListenAndServe())
}
