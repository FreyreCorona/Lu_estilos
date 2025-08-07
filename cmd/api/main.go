package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/FreyreCorona/Lu_estilos/internal/models"
	_ "github.com/lib/pq"
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	var cfg configuration
	// load the environment variables
	flag.IntVar(&cfg.port, "port", 4000, "Port number")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://user:password@addres:port/db_name", "DSN for postgres database")

	flag.Parse()

	Infolog := log.New(os.Stdout, "", log.Ltime|log.Ldate)

	db, err := openDB(cfg.db.dsn)
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
