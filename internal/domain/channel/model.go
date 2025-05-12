package channel

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model

	Name         string `gorm:"unique;not null"`
	IsDirect     bool
	Description         *string

	Participants []ChannelParticipant `gorm:"many2many:channel_participants;"`
	ImageID      *uint
}

type ChannelParticipant struct {
	gorm.Model

	UserID   uint
	Role  ChannelParticipantRole `gorm:"type:enum('creator','admin','participant');not null"`
	Channels []Channel `gorm:"many2many:channel_participants;"`
}
