package channel

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model

	Name         string
	IsDirect     bool
	Description         *string

	Participants []ChannelParticipant `gorm:"many2many:channel_participants;"`
	ImageID      *uint
}

type ChannelParticipant struct {
	gorm.Model

	UserID   uint
	Channels []Channel `gorm:"many2many:channel_participants;"`
}
