package file

import (
	"github.com/darksuei/chat-kit/internal/domain/channel"
	"github.com/darksuei/chat-kit/internal/domain/message"
	"gorm.io/gorm"
)

type File struct {
    gorm.Model
	Identifier    string
	Name          string

    ChannelID     uint
    Channel       channel.Channel `gorm:"foreignKey:ChannelID"`

    ParticipantID uint
    Participant   channel.ChannelParticipant `gorm:"foreignKey:ParticipantID"`

	MessageID     *uint
	Message       *message.Message `gorm:"foreignKey:MessageID"`

	VersionID     uint
	Versions      []FileVersion `gorm:"foreignKey:FileID"`

	PermissionID  *uint
	Scope         *string
}

type FileVersion struct {
	gorm.Model

	FileID    uint
	File      File `gorm:"foreignKey:FileID"`

	Revision  string `gorm:"default:'LATEST'"`
	IsActive  bool   `gorm:"default:true"`
}
