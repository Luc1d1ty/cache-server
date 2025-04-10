package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/Luc1d1ty/cache-server/internal/api"   
	"github.com/Luc1d1ty/cache-server/internal/cache" 
)

func main() {
    port := flag.String("port", "8080", "HTTP server port")
    flag.Parse()
    
    cacheInstance := cache.NewCache()

    quit := make(chan bool)
    cacheInstance.StartTTLManager(10*time.Second, quit)

    apiHandler := &api.APIHandler{Cache: cacheInstance}

    mux := http.NewServeMux()
    mux.HandleFunc("/cache/set", apiHandler.SetHandler)
    mux.HandleFunc("/cache/get", apiHandler.GetHandler)
    mux.HandleFunc("/cache/delete", apiHandler.DeleteHandler)
    mux.HandleFunc("/cache/metrics", apiHandler.MetricsHandler)

    loggedMux := api.LoggingMiddleware(mux)

    log.Printf("Starting HTTP server on port %s...\n", *port)
    if err := http.ListenAndServe(":"+*port, loggedMux); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }

    // quit <- true
}