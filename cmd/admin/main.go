package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"service":    "admin",
			"goroutines": runtime.NumGoroutine(),
			"gomaxprocs": runtime.GOMAXPROCS(0),
			"memlimit":   os.Getenv("GOMEMLIMIT"),
		})
	})
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}
