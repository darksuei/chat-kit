package channel

import (
	database "github.com/darksuei/chat-kit/internal/infrastructure/database"
)

type Service interface {
	CreateChannel(payload *ChannelInterface) (*Channel, error)
	UpdateChannel(id int64, payload *OptionalChannelInterface) error
	GetChannels(where *OptionalChannelInterface) (*[]Channel, error)
	GetChannelById(id int64) (*Channel, error)

	CreateChannelParticipant(userId string, channelId uint, role ChannelParticipantRole) error
	DeleteChannelParticipant(userId string, channelId uint) error
	FindChannelParticipant(userId string, channelId uint) (*ChannelParticipant, error)
	GetListOfParticipants(userIdList *[]string, channelId uint) (*[]ChannelParticipant, error)
}
type serviceDefinition struct{}

var repo Repository = NewRepository()

func (s *serviceDefinition) CreateChannel(payload *ChannelInterface) (*Channel, error) {
	return repo.Create(database.DB, payload)
}

func (s *serviceDefinition) UpdateChannel(id int64, payload *OptionalChannelInterface) error {
	channel, err := repo.FindById(database.DB, id)

	if err != nil {
		return err
	}

	if payload.Name != nil {
		channel.Name = *payload.Name
	}

	if payload.IsDirect != nil {
		channel.IsDirect = *payload.IsDirect
	}

	if payload.Description != nil {
		channel.Description = payload.Description
	}

	err = repo.Update(database.DB.Model(&channel), channel)

	return err
}

func (s *serviceDefinition) GetChannels(where *OptionalChannelInterface) (*[]Channel, error) {
	return repo.Find(database.DB, where)
}

func (s *serviceDefinition) GetChannelById(id int64) (*Channel, error) {
	return repo.FindById(database.DB, id)
}

func (s *serviceDefinition) CreateChannelParticipant(userId string, channelId uint, role ChannelParticipantRole) error {
	return repo.CreateParticipant(database.DB, userId, channelId, role)
}

func (s *serviceDefinition) DeleteChannelParticipant(userId string, channelId uint) error {
	return repo.DeleteParticipant(database.DB, userId, channelId)
}

func (s *serviceDefinition) FindChannelParticipant(userId string, channelId uint) (*ChannelParticipant, error) {
	return repo.FindParticipant(database.DB, userId, channelId)
}

func (s *serviceDefinition) GetListOfParticipants(userIdList *[]string, channelId uint) (*[]ChannelParticipant, error) {
	return repo.GetListOfParticipants(database.DB, userIdList, channelId)
}

func NewService() Service {
	return &serviceDefinition{}
}