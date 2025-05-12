package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetHeader("x-user-id")
		
		if userId == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized."})
			return
		}
		
		log.Printf("Retrieved user id %s", userId)

		// Set user ID in context
		c.Set("userId", userId)
		
		c.Next()
	}
}