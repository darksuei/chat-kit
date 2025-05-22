package message

import (
	"gorm.io/gorm"

	channel "github.com/darksuei/chat-kit/internal/domain/channel"
)

type Message struct {
	gorm.Model

	ChannelID     uint
	Channel       channel.Channel `gorm:"constraint:OnDelete:CASCADE;"`

	ParticipantID uint
	Participant   channel.ChannelParticipant `gorm:"constraint:OnDelete:SET NULL;"`

	Content       string `gorm:"type:text;not null"`

	ReadBy    []channel.ChannelParticipant `gorm:"many2many:message_read_by;"`
	Mentions  []channel.ChannelParticipant `gorm:"many2many:message_mentions;"`

	Reactions []MessageReaction `gorm:"constraint:OnDelete:CASCADE;"`

	IsChild  bool   `gorm:"default:false"`
	ParentID *uint
	Parent   *Message `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;"`
}


type MessageReaction struct {
	gorm.Model

	MessageID uint    
	Message   Message

	Emoji     string
}