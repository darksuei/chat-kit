package api

import (
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	log.Print("Received health check request.")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is healthy!"))
}