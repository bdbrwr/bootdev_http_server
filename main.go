package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1) // Increment the counter
		next.ServeHTTP(w, r)      // Call the next handler in the chain
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	currentHits := cfg.fileserverHits.Load()
	w.Write([]byte(fmt.Sprintf("Hits: %d", currentHits)))
}

func main() {
	const port = "8080"
	const filepathRoot = "."

	apiConfig := apiConfig{}
	apiConfig.fileserverHits.Store(0)

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", apiConfig.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("/healthz", handlerReadiness)
	mux.HandleFunc("/metrics", apiConfig.handlerMetrics)
	mux.HandleFunc("/reset", apiConfig.handlerReset)

	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
