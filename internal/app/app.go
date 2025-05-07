package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/darksuei/chat-kit/config"
	"github.com/darksuei/chat-kit/internal/api"
	"github.com/darksuei/chat-kit/internal/database"

	"github.com/joho/godotenv"
)

func Run() {
	godotenv.Load()

	port, err := config.ReadEnv("PORT")

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", api.Health)

	server := &http.Server{Addr: ":" + port}

	go func() {
		log.Printf("Application is running on port: %s...", port)

		database.Connect()
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %s", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop // Wait for the signal to terminate

	log.Println("Shutting down application...")

	// Gracefully shutdown HTTP server
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Application shutdown failed: %s...", err)
	}

	log.Println("Application shutdown...")
}