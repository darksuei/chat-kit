package message

type MessageInterface struct {
	UserID    string   `json:"user_id" binding:"required"`
	Content   string   `json:"content"`
	Mentions  *[]string `json:"mentions"`
	IsChild   *bool    `json:"is_child"`
	ParentID  *uint    `json:"parent_id"`
}

type RecentMessagesPayload struct {
	Type     string              `json:"type"`
	Messages *[]Message  `json:"messages"`
}

type RawMessage struct {
	Type     string              `json:"type" binding:"required"`
	Payload interface{}  `json:"payload" binding:"required"`
}