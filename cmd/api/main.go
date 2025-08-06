package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
)

type configuration struct {
	port int
}
type application struct {
	Config     configuration
	InfoLogger *log.Logger
	Clients    *models.ClietModel
}

func main() {
	var cfg configuration
	// set flago to custom port number
	flag.IntVar(&cfg.port, "port", 4000, "Port listen number")
	flag.Parse()

	Infolog := log.New(os.Stdout, "", log.Ltime|log.Ldate)

	// initialize application struct
	app := application{
		Config:     cfg,
		InfoLogger: Infolog,
	}
	// initialize server struct
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	Infolog.Printf("Starting Server at %s ...\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
