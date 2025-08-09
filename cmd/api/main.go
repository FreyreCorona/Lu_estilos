package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FreyreCorona/Lu_estilos/internal/models"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type configuration struct {
	port int
	db   struct {
		dsn string
	}
}

type application struct {
	Config  configuration
	Logger  *log.Logger
	Clients *models.ClietModel
}

func main() {
	var cfg configuration

	// load the environment variables
	cfg.port = GetEnvInt("API_PORT", 4000)
	cfg.db.dsn = GetEnvStr("DSN", " ")

	Infolog := log.New(os.Stdout, "", log.Ltime|log.Ldate)

	db, err := openDB(cfg.db.dsn)
	if err != nil {
		Infolog.Fatal(err)
	}
	defer db.Close()

	// initialize application struct
	app := application{
		Config:  cfg,
		Logger:  Infolog,
		Clients: &models.ClietModel{DB: db},
	}
	err = app.ApplicateMigrations("file:///app/migrations")
	if err != nil {
		Infolog.Fatal(err)
	}
	// initialize server struct
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	Infolog.Printf("Starting Server at port: %s ...\n", server.Addr)
	Infolog.Fatal(server.ListenAndServe())
}
