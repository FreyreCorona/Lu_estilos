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
	db   struct {
		dsn string
	}
}

type application struct {
	Config     configuration
	InfoLogger *log.Logger
	Clients    *models.ClietModel
}

func main() {
	cfg := configuration{
		port: GetEnvInt("API_PORT", 4000),
	}
	// load the environment variables
	flag.IntVar(&cfg.port, "port", GetEnvInt("API_PORT", 4000), "Port number")
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("DSN"), "DSN for postgres database")

	flag.Parse()

	Infolog := log.New(os.Stdout, "", log.Ltime|log.Ldate)

	db, err := openDB(cfg.db.dsn)
	if err != nil {
		Infolog.Fatal(err)
	}
	defer db.Close()

	if err != nil {
		Infolog.Fatal(err)
	}

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
