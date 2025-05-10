package app

import (
	"github.com/gin-gonic/gin"

	"github.com/darksuei/chat-kit/internal/api"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/health", api.Health)

	return router
}