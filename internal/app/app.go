package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/darksuei/chat-kit/config"
	"github.com/darksuei/chat-kit/internal/database"
	"github.com/joho/godotenv"
)

func Run() {
	// Load environment variables
	_ = godotenv.Load()

	port, err := config.ReadEnv("PORT")
	if err != nil {
		log.Fatal("Error reading PORT from environment:", err)
	}

	database.Connect()

	router := Router()

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		log.Printf("Application is running on port: %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Application startup failed: %s", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down application..")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %s", err)
	}

	log.Println("Application shutdown successfully..")
}
