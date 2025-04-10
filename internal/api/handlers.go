package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Luc1d1ty/cache-server/internal/cache"
)

type APIHandler struct {
	Cache *cache.Cache
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   int    `json:"ttl,omitempty"`
}

func (h *APIHandler) SetHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("[SetHandler] Received a /cache/set request")

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("[SetHandler] Error reading request body: %v", err)
        http.Error(w, "Unable to read request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    var req SetRequest
    err = json.Unmarshal(body, &req)
    if err != nil {
        log.Printf("[SetHandler] Invalid JSON: %v", err)
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    if req.Key == "" {
        log.Println("[SetHandler] Missing key in request")
        http.Error(w, "Key is required", http.StatusBadRequest)
        return
    }

    h.Cache.Set(req.Key, req.Value, req.TTL)
    log.Printf("[SetHandler] Key '%s' set successfully", req.Key)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Key set successfully"))
}

func (h *APIHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is required", http.StatusBadRequest)
		return
	}

	value, ok := h.Cache.Get(key)
	if !ok {
		http.Error(w, "Key not found or expired", http.StatusNotFound)
		return
	}

	resp := map[string]string{"key": key, "value": value}
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func (h *APIHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is required", http.StatusBadRequest)
		return
	}

	err := h.Cache.Delete(key)
	if err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Key deleted successfully"))
}

func (h *APIHandler) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := h.Cache.GetMetrics()
	jsonResp, _ := json.Marshal(metrics)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
