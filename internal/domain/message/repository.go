package message

import (
	"errors"

	"gorm.io/gorm"

	channel "github.com/darksuei/chat-kit/internal/domain/channel"
)

type Repository interface {
	Create(db *gorm.DB, payload *MessageInterface, participant *channel.ChannelParticipant, mentions *[]channel.ChannelParticipant, channel *channel.Channel) (*Message, error)
	GetMessages(db *gorm.DB, channel *channel.Channel, limit *int, beforeMessageId *uint64) (*[]Message, error)
}

type repositoryDefinition struct {}

func (r *repositoryDefinition) Create(db *gorm.DB, payload *MessageInterface, participant *channel.ChannelParticipant, mentions *[]channel.ChannelParticipant, channel *channel.Channel) (*Message, error) {
	message := Message{ChannelID: channel.ID, ParticipantID: participant.ID, Content: payload.Content, Mentions: *mentions, IsChild: *payload.IsChild, ParentID: payload.ParentID}

	err := db.Create(&message).Error
	if err != nil {
		return nil, errors.New("failed to save message")
	}

	err = db.Preload("Participant").Preload("Mentions").First(&message, message.ID).Error
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (r *repositoryDefinition) GetMessages(db *gorm.DB, channel *channel.Channel, limit *int, beforeMessageId *uint64) (*[]Message, error) {
	var messages []Message

	if limit == nil || *limit == 0 {
		defaultLimit := 100
		limit = &defaultLimit
	}

	// Start building query
	query := db.
		Where("channel_id = ?", channel.ID).
		Order("created_at DESC").
		Limit(*limit)

	// If beforeMessageId is set, fetch the message and use its timestamp
	if beforeMessageId != nil && *beforeMessageId != 0 {
		var referenceMsg Message
		if err := db.Select("created_at").First(&referenceMsg, *beforeMessageId).Error; err == nil {
			query = query.Where("created_at < ?", referenceMsg.CreatedAt)
		} else {
			return nil, err
		}
	}

	// Execute query
	if err := query.Find(&messages).Error; err != nil {
		return nil, err
	}

	return &messages, nil
}



func NewRepository() Repository {
	return &repositoryDefinition{}
}