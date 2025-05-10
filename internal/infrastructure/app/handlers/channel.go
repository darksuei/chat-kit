package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	channel "github.com/darksuei/chat-kit/internal/domain/channel"
	helpers "github.com/darksuei/chat-kit/internal/helpers"
	database "github.com/darksuei/chat-kit/internal/infrastructure/database"
)

var repository channel.Repository = channel.NewRepository()

func CreateChannel(c *gin.Context) {
	payload, err := helpers.ValidateRequest[channel.ChannelInterface](c)

	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	err = repository.Create(database.DB, payload)

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to create channel: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Channel created successfully."})
}

func GetChannels(c *gin.Context) {
	where := channel.OptionalChannelInterface{}

	channels, err := repository.Find(database.DB, &where)

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch channels: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"channels": channels})
}

func GetChannelById(c *gin.Context) {
	id := c.Param("id")

	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	channel, err := repository.FindById(database.DB, idInt64)

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to fetch channel: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"channel": channel})
}