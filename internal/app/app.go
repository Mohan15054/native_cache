package app

import (
	"cachesystem/api"
	"cachesystem/internal/config"
	"fmt"
)

func Run() {
	cfg := config.Load()
	fmt.Println("Cache system starting with config:", cfg)

	go api.StartServer(cfg.ListenAddr)
	select {} // Block forever
}

// To access the cache system's REST API after starting the server on port 8080:
//
// Set a key:
//   curl "http://localhost:8080/set?key=mykey&value=myvalue&ttl=60"
//
// Get a key:
//   curl "http://localhost:8080/get?key=mykey"
//
// Delete a key:
//   curl "http://localhost:8080/delete?key=mykey"
//
// Get metrics:
//   curl "http://localhost:8080/metrics"
//
// Replace 'mykey' and 'myvalue' with your desired key and value.
// The 'ttl' parameter is the time-to-live in seconds.
