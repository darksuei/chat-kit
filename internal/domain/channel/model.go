package channel

import (
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model

	Name         string `gorm:"unique;not null"`
	IsDirect     bool
	Description         *string

	Participants []ChannelParticipant
	ImageID      *uint
}

type ChannelParticipant struct {
	gorm.Model

	UserID   string
	Role  ChannelParticipantRole `gorm:"type:enum('creator','admin','participant');not null"`

	ChannelID uint
}
