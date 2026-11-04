package main

import (
	"log/slog"
	"os"
	"time"
)

// A stand-in for a queue consumer: it "processes" a job every 5 seconds.
func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	log.Info("worker started")
	for n := 1; ; n++ {
		time.Sleep(5 * time.Second)
		log.Info("processed job", "job", n)
	}
}
