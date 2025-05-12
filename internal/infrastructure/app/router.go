package app

import (
	"github.com/gin-gonic/gin"

	handlers "github.com/darksuei/chat-kit/internal/infrastructure/app/handlers"
	middlewares "github.com/darksuei/chat-kit/internal/infrastructure/app/middlewares"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.Health)

	/*
	* Channels
	*/
	router.GET("/channel", middlewares.AuthMiddleware(), handlers.GetChannels)
	router.GET("/channel/:id", middlewares.AuthMiddleware(), handlers.GetChannelById)
	router.PUT("/channel", middlewares.AuthMiddleware(), handlers.UpdateChannel)
	router.POST("/channel", middlewares.AuthMiddleware(), handlers.CreateChannel)

	return router
}