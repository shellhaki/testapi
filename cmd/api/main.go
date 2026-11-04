package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var started = time.Now()

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if os.Getenv("DATABASE_URL") == "" {
		panic("DATABASE_URL is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "hello"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		reply(w, map[string]any{"message": "api running", "version": 7})
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		reply(w, map[string]any{"message": "alive", "up_for": time.Since(started).Round(time.Second).String()})
	})
	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		log.Info("hello", "name", r.PathValue("name"))
		reply(w, map[string]any{"message": greeting + ", " + r.PathValue("name")})
	})

	log.Info("server started", "port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func reply(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
