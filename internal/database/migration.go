package database

import (
	"log"

	"gorm.io/gorm"

	"github.com/darksuei/chat-kit/internal/domain/channel"
	"github.com/darksuei/chat-kit/internal/domain/file"
	"github.com/darksuei/chat-kit/internal/domain/message"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&channel.Channel{},
		&channel.ChannelParticipant{},
		&message.Message{},
		&message.MessageReaction{},
		&file.File{},
		&file.FileVersion{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully..")
}