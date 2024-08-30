package main

import (
	"blog_api/internal/router"
	"blog_api/internal/storage"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load("config/.env"); err != nil {
		return err
	}
	return nil
}

func main() {
	// Load environment variables
	err := LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database
	storage.InitDB()
	r := router.SetUpRouter()

	// Set up the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Print("Server is starting at port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe Failed: %v", err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Println("Received interruption signal, Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown Failed: %v", err)
	}

	log.Println("Server exited gracefully")
}
