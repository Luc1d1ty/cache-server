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

	http.HandleFunc("/cache/set", apiHandler.SetHandler)
	http.HandleFunc("/cache/get", apiHandler.GetHandler)
	http.HandleFunc("/cache/delete", apiHandler.DeleteHandler)
	http.HandleFunc("/cache/metrics", apiHandler.MetricsHandler)

	log.Printf("Starting HTTP server on port %s...\n", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// quit <- true
}
