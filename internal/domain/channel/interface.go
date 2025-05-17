package channel

type ChannelInterface struct {
	Name        string `json:"name" binding:"required"`
	IsDirect    *bool   `json:"is_direct" binding:"required"`
	Description string `json:"description"`
}

type OptionalChannelInterface struct {
	Name        *string `json:"name"`
	IsDirect    *bool   `json:"is_direct"`
	Description *string `json:"description"`
}

type ChannelParticipantInterface struct {
	UserID string `json:"user_id" binding:"required"`
	// Role ChannelParticipantRole `json:"role" binding:"required"`
	ChannelID uint `json:"channel_id" binding:"required"`
}