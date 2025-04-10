package api

import (
    "log"
    "net/http"
    "time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("[Request] %s %s started", r.Method, r.RequestURI)
        next.ServeHTTP(w, r)
        duration := time.Since(start)
        log.Printf("[Request] %s %s completed in %v", r.Method, r.RequestURI, duration)
    })
}
