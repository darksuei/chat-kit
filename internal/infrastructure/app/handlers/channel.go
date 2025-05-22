package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	channel "github.com/darksuei/chat-kit/internal/domain/channel"
	helpers "github.com/darksuei/chat-kit/internal/helpers"
)

var service channel.Service = channel.NewService()

func CreateChannel(c *gin.Context) {
	payload, err := helpers.ValidateRequest[channel.ChannelInterface](c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	log.Println("Creating channel..")

	newChannel, err := service.CreateChannel(payload)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create channel: " + err.Error()})
		return
	}

	log.Println("Channel: ", newChannel)

	log.Println("Creating channel participant..")

	userId, err := helpers.GetUserIdFromContext(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user: " + err.Error()})
		return
	}

	err = service.CreateChannelParticipant(*userId, newChannel.ID, channel.Creator)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create channel participant: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Channel created successfully."})
}

func UpdateChannel(c *gin.Context) {
	payload, err := helpers.ValidateRequest[channel.OptionalChannelInterface](c)
	
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	id := c.Param("id")

	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = service.UpdateChannel(idInt64, payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update channel: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Channel updated successfully."})
}

func GetChannels(c *gin.Context) {
	where := channel.OptionalChannelInterface{}

	channels, err := service.GetChannels(&where)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch channels: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"channels": channels})
}

func GetChannelById(c *gin.Context) {
	id := c.Param("id")

	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	channel, err := service.GetChannelById(idInt64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch channel: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"channel": channel})
}

func CreateChannelParticipant(c *gin.Context) {
	payload, err := helpers.ValidateRequest[channel.ChannelParticipantInterface](c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	log.Println("Checking channel..")

	existingChannel, err := service.GetChannelById(int64(payload.ChannelID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Channel does not exist: " + err.Error()})
		return
	}

	log.Println("Checking duplicate participant..")

	_, err = service.FindChannelParticipant(payload.UserID, existingChannel.ID)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Participant already exists."})
		return
	}

	log.Println("Checking admin participant..")

	userId, err := helpers.GetUserIdFromContext(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user: " + err.Error()})
		return
	}

	log.Print(*userId, existingChannel.ID)

	admin, err := service.FindChannelParticipant(*userId, existingChannel.ID)
	
	if err != nil || admin.Role != channel.Creator && admin.Role != channel.Admin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient permissions to perform this action: " + err.Error()})
		return
	}

	log.Println("Adding channel participant..")

	err = service.CreateChannelParticipant(payload.UserID, payload.ChannelID, channel.Participant)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add channel participant: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Channel participant added successfully."})
}

func RemoveChannelParticipant(c *gin.Context) {
	payload, err := helpers.ValidateRequest[channel.ChannelParticipantInterface](c)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	log.Println("Checking channel..")

	existingChannel, err := service.GetChannelById(int64(payload.ChannelID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Channel does not exist: " + err.Error()})
		return
	}

	log.Println("Checking admin participant..")

	userId, err := helpers.GetUserIdFromContext(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user: " + err.Error()})
		return
	}

	log.Print(*userId, existingChannel.ID)

	admin, err := service.FindChannelParticipant(*userId, existingChannel.ID)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient permissions to perform this action: " + err.Error()})
		return
	}

	existingParticipant, err := service.FindChannelParticipant(payload.UserID, existingChannel.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid participant: " + err.Error()})
		return
	}

	if existingParticipant.Role == channel.Creator {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot remove the creator of this channel."})
		return
	}
	
	if admin.Role == channel.Participant || admin.Role == channel.Admin && existingParticipant.Role == channel.Admin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient permissions to perform this action."})
		return
	}

	log.Println("Removing channel participant..")

	err = service.DeleteChannelParticipant(payload.UserID, payload.ChannelID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to remove channel participant: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Channel participant removed successfully."})
}