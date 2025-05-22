package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	channel "github.com/darksuei/chat-kit/internal/domain/channel"
	message "github.com/darksuei/chat-kit/internal/domain/message"
)

var messageService message.Service = message.NewService()
var channelService channel.Service = channel.NewService()

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Todo: Handle origin check
		return true
	},
}

func HandleMessageWebsocket(c *gin.Context) {
	channelIdStr := c.Param("id")

	channelId, err := strconv.ParseInt(channelIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	messageChannel, err := channelService.GetChannelById(channelId)
	if err != nil {
		log.Println("Failed to get connection channel: ", err)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println("Failed to upgrade websocket: ", err)
		return
	}

	log.Println("Client connected to socket..", conn.RemoteAddr())

	// 01. When the client connects, send back the most recent messages if any.
	recentMessages, err := messageService.GetMessages(messageChannel, nil, nil)
	if err != nil {
		log.Println("failed to retreive recent messages:", err)
		return
	}

	recentMessagesPayload := message.RecentMessagesPayload{
		Type: string(message.RecentMessages),
		Messages: recentMessages,
	}

	recentMessagesBytes, err := json.Marshal(recentMessagesPayload)
	if err != nil {
		fmt.Println("failed to marshal message: ", err)
	}
	
	err = conn.WriteMessage(websocket.TextMessage, recentMessagesBytes)
	if err != nil {
		log.Println("Write failed:", err)
		return
	}

	defer conn.Close()

	// Keep listening for messages
	for {
		messageType, rawMessage, err := conn.ReadMessage()
		if err != nil {
			break // connection closed or error occurred
		}

		log.Println("Message type: ", messageType)
		log.Println("Message: ", string(rawMessage))

		parsedMessage, err := messageService.ParseRawMessage(rawMessage)
		if err != nil {
			fmt.Println("failed to parse message: ", err)
			continue
		}

		switch parsedMessage.Type {
			case string(message.PublishedMessage):
				payload, ok := parsedMessage.Payload.(message.MessageInterface)
				if !ok {
					log.Println("Invalid message type")
					continue
				}

				savedMessage, err := messageService.CreateMessage(payload, messageChannel)
				if err != nil {
					fmt.Println("failed to save message: ", err)
					continue
				}

				savedMessageBytes, err := json.Marshal(savedMessage)
				if err != nil {
					fmt.Println("failed to marshal message: ", err)
				}
		
				if err := conn.WriteMessage(messageType, savedMessageBytes); err != nil {
					return
				}
			
			case string(message.ReadMessages):
				// handle a read messages event
				// takes an array of messages and marks them as read for the given user
		}
	}
}

// Mark message as read
// Fix mentions
// Allow child messages
// Allow reactions to message
// Fix saved channel bug
func GetMessages(c *gin.Context){
	channelIdStr := c.Param("id")

	channelId, err := strconv.ParseInt(channelIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "100")

	limit, err := strconv.Atoi(limitStr)
    if err != nil {
        log.Println("error parsing limit, skipping...")
    }

	beforeMessageIdStr := c.Query("beforeMessageId")

	beforeMessageId, err := strconv.ParseUint(beforeMessageIdStr, 10, 64)
    if err != nil {
        log.Println("beforeMessageId not provided, skipping...")
    }

	messageChannel, err := channelService.GetChannelById(channelId)
	if err != nil {
		log.Println("Failed to get connection channel: ", err)
		return
	}
	
	messages, err := messageService.GetMessages(messageChannel, &limit, &beforeMessageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}