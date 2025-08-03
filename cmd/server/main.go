package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FreyreCorona/Lu_estilos/api"
)

func main() {
	Infolog := log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate)
	Errorlog := log.New(os.Stdout, "ERROR: ", log.Ltime|log.Ldate|log.Lshortfile)
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.Handle)

	server := http.Server{
		Addr:     ":8000",
		Handler:  mux,
		ErrorLog: Errorlog,
	}
	Infolog.Println("Starting Server...")
	Errorlog.Fatal(server.ListenAndServe())
}
