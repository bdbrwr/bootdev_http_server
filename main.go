package main

import (
	"log"
	"net/http"
)

const port = "8080"

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
