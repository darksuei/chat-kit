package message

type MessageType string

const (
	RecentMessages   MessageType = "recent_messages"
	ReadMessages   MessageType = "read_messages"
	PublishedMessage   MessageType = "published_message"
)