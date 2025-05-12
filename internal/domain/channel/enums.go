package channel

type ChannelParticipantRole string

const (
	Creator   ChannelParticipantRole = "creator"
	Admin  ChannelParticipantRole = "admin"
	Participant ChannelParticipantRole = "participant"
)
