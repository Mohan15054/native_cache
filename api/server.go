package api

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

var (
	cache   = make(map[string]item)
	cacheMu sync.RWMutex
)

type item struct {
	Value      string
	ExpiryTime time.Time
}

func StartServer(addr string) {
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(addr, nil)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	ttlStr := r.URL.Query().Get("ttl")
	if key == "" || value == "" || ttlStr == "" {
		http.Error(w, "Missing key, value, or ttl", http.StatusBadRequest)
		return
	}

	ttl, err := time.ParseDuration(ttlStr + "s")
	if err != nil {
		http.Error(w, "Invalid ttl", http.StatusBadRequest)
		return
	}

	cacheMu.Lock()
	cache[key] = item{Value: value, ExpiryTime: time.Now().Add(ttl)}
	cacheMu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Key set successfully"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	cacheMu.RLock()
	it, exists := cache[key]
	cacheMu.RUnlock()

	if !exists || time.Now().After(it.ExpiryTime) {
		http.Error(w, "Key not found or expired", http.StatusNotFound)
		return
	}

	response := map[string]string{"value": it.Value}
	json.NewEncoder(w).Encode(response)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Missing key", http.StatusBadRequest)
		return
	}

	cacheMu.Lock()
	delete(cache, key)
	cacheMu.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Key deleted successfully"))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	cacheMu.RLock()
	totalKeys := len(cache)
	cacheMu.RUnlock()

	response := map[string]int{"total_keys": totalKeys}
	json.NewEncoder(w).Encode(response)
}
