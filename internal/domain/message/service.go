package message

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"

	channel "github.com/darksuei/chat-kit/internal/domain/channel"
	database "github.com/darksuei/chat-kit/internal/infrastructure/database"
)

type Service interface {
	ParseRawMessage(raw []byte) (*RawMessage, error)
	CreateMessage(payload MessageInterface, channel *channel.Channel) (*Message, error)
	GetMessages(channel *channel.Channel, limit *int, beforeMessageId *uint64) (*[]Message, error)
}

type serviceDefinition struct{}

var messageRepository Repository = NewRepository()

var channelService channel.Service = channel.NewService()

func (s *serviceDefinition) ParseRawMessage(raw []byte) (*RawMessage, error) {
	var parsedMessage RawMessage

	log.Println("Parsing received raw message..")

	log.Println(parsedMessage)

	err := json.Unmarshal(raw, &parsedMessage)
    if err != nil {
        fmt.Println("Error parsing message: ", err)
        return nil, err
    }

	log.Println(parsedMessage)

	validate := validator.New()
	if err := validate.Struct(parsedMessage); err != nil {
		return nil, err
	}

	return &parsedMessage, nil
}

func (s *serviceDefinition) CreateMessage(payload MessageInterface, channel *channel.Channel) (*Message, error) {
	participant, err := channelService.FindChannelParticipant(payload.UserID, channel.ID)
	if err != nil {
		return nil, errors.New("invalid participant")
	}

	mentions, err := channelService.GetListOfParticipants(payload.Mentions, channel.ID)
	if err != nil {
		fmt.Println("Error fetching mentions: ", err)
		return nil, err
	}

	// Save message to the database
	message, err := messageRepository.Create(database.DB, &payload, participant, mentions, channel)
	if err != nil {
		fmt.Println("Error saving message: ", err)
		return nil, err
	}

	return message, nil
}

func (s *serviceDefinition) GetMessages(channel *channel.Channel, limit *int, beforeMessageId *uint64) (*[]Message, error) {
	return messageRepository.GetMessages(database.DB, channel, limit, beforeMessageId)
}

func NewService() Service {
	return &serviceDefinition{}
}