package channel

import (
	database "github.com/darksuei/chat-kit/internal/infrastructure/database"
)

type Service interface {
	CreateChannel(payload *ChannelInterface) error
}

type serviceDefinition struct{}

var repo Repository = NewRepository()

func CreateChannel(payload *ChannelInterface) error {
	return repo.Create(database.DB, payload)
}

func CreateChannelParticipant() error {
	// given the created channel and the userId, add the participant as the creator
	
}