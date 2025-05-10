package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	log.Print("Received health check request.")
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is healthy!",
	  })
}