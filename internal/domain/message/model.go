package message

import (
	"gorm.io/gorm"

	channel "github.com/darksuei/chat-kit/internal/domain/channel"
)

type Message struct {
    gorm.Model

    ChannelID     uint
    Channel       channel.Channel

    ParticipantID uint
    Participant   channel.ChannelParticipant

    Content       string

    ReadBy        []channel.ChannelParticipant `gorm:"many2many:message_read_by;"`
    Mentions      []channel.ChannelParticipant `gorm:"many2many:message_mentions;"`
    Reactions     []MessageReaction

    IsChild       bool `gorm:"default:false"`
    ParentID      *uint
    Parent        *Message
}

type MessageReaction struct {
	gorm.Model

	MessageID uint    
	Message   Message

	Emoji     string
}