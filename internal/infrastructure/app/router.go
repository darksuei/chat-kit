package app

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/darksuei/chat-kit/internal/infrastructure/app/handlers"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.Health)

	/*
	* Channels
	*/
	router.GET("/channel", handlers.GetChannels)
	router.GET("/channel/:id", handlers.GetChannelById)
	router.POST("/channel", handlers.CreateChannel)

	return router
}