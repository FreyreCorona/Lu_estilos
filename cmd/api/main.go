package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
	"github.com/joho/godotenv"
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
	// load the environment variables
	godotenv.Load(".env")
	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		port = 4000
	}
	cfg.port = port

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
