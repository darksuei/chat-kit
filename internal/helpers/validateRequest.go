package helpers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ValidateRequest[T any](c *gin.Context) (*T, error) {
	var payload T

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Printf("Invalid request body: %s", err.Error())
		return nil, err
	}

	return &payload, nil
}
