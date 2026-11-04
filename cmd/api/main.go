package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var started = time.Now()

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if os.Getenv("DATABASE_URL") == "" {
		panic("DATABASE_URL is not set")
	}
	// Parsed only: there's no database behind this test, but the driver is linked in like a real app's.
	if _, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL")); err != nil {
		log.Warn("bad DATABASE_URL", "err", err)
	}
	greeting := os.Getenv("GREETING")
	if greeting == "" {
		greeting = "hello"
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "api running", "version": 11, "framework": "gin"}) })
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "alive", "up_for": time.Since(started).Round(time.Second).String()}) })
	r.GET("/hello/:name", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": greeting + ", " + c.Param("name")}) })
	log.Info("server started", "port", os.Getenv("PORT"))
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}
}
