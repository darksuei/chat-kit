package channel

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindById(db *gorm.DB, id int64) (*Channel, error)
	FindOne(db *gorm.DB, where *OptionalChannelInterface) (*Channel, error)
	Find(db *gorm.DB, where *OptionalChannelInterface) (*[]Channel, error)
	Create(db *gorm.DB, payload *ChannelInterface) (*Channel, error)
	Update(db *gorm.DB, channel *Channel) error

	CreateParticipant(db *gorm.DB, userId string, channelId uint, role ChannelParticipantRole) error
	DeleteParticipant(db *gorm.DB, userId string, channelId uint) error
	FindParticipant(db *gorm.DB, userId string, channelId uint) (*ChannelParticipant, error)
	GetListOfParticipants(db *gorm.DB, userIdList *[]string, channelId uint) (*[]ChannelParticipant, error)
}

type repositoryDefinition struct{}

func (r *repositoryDefinition) FindById(db *gorm.DB, id int64) (*Channel, error) {
	var channel Channel

	if err := db.Unscoped().First(&channel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Channel not found")
		}
		return nil, err
	}

	return &channel, nil
}

func (r *repositoryDefinition) FindOne(db *gorm.DB, where *OptionalChannelInterface) (*Channel, error) {
	query := map[string]interface{}{}

	if where.Name != nil {
		query["name"] = *where.Name
	}
	if where.Description != nil {
		query["description"] = *where.Description
	}
	if where.IsDirect != nil {
		query["is_direct"] = *where.IsDirect
	}
	
	var channel Channel

	if err := db.Unscoped().Where(query).First(&channel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Channel not found")
		}
		return nil, err
	}

	return &channel, nil	
}

func (r *repositoryDefinition) Find(db *gorm.DB, where *OptionalChannelInterface) (*[]Channel, error) {
	query := map[string]interface{}{}

	if where.Name != nil {
		query["name"] = *where.Name
	}
	if where.Description != nil {
		query["description"] = *where.Description
	}
	if where.IsDirect != nil {
		query["is_direct"] = *where.IsDirect
	}
	
	var channels []Channel

	if err := db.Unscoped().Preload("Participants").Where(query).Find(&channels).Error; err != nil {
		return nil, err
	}

	return &channels, nil	
}

func (r *repositoryDefinition) Create(db *gorm.DB, payload *ChannelInterface) (*Channel, error) {
	channel := Channel{Name: payload.Name, IsDirect: *payload.IsDirect, Description: &payload.Description}

	err := db.Create(&channel).Error

	if err != nil {
		return nil, errors.New("failed to create channel: " + err.Error())
	}

	return &channel, nil
}

func (r *repositoryDefinition) Update(db *gorm.DB, channel *Channel) error {
	err := db.Updates(channel).Error

	if err != nil {
		return errors.New("failed to update channel: " + err.Error())
	}

	return nil
}

func (r *repositoryDefinition) CreateParticipant(db *gorm.DB, userId string, channelId uint, role ChannelParticipantRole) error {
	participant := ChannelParticipant{
		UserID:  userId,
		Role:    role,
		ChannelID: channelId,
	}

	if err := db.Create(&participant).Error; err != nil {
		return err
	}

	return nil
}

func (r *repositoryDefinition) DeleteParticipant(db *gorm.DB, userId string, channelId uint) error {
	var participant ChannelParticipant

	// First, find the participant
	err := db.Where("user_id = ? AND channel_id = ?", userId, channelId).
		First(&participant).Error

	if err != nil {
		return err
	}

	// Then delete the participant
	if err := db.Delete(&participant).Error; err != nil {
		return err
	}

	return nil
}

func (r *repositoryDefinition) FindParticipant(db *gorm.DB, userId string, channelId uint) (*ChannelParticipant, error) {
	var participant ChannelParticipant

	err := db.
		Unscoped().
		Where("user_id = ? AND channel_id = ?", userId, channelId).
		First(&participant).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("participant not found in the specified channel")
		}
		return nil, err
	}

	return &participant, nil
}

func (r * repositoryDefinition) GetListOfParticipants(db *gorm.DB, userIdList *[]string, channelId uint) (*[]ChannelParticipant, error) {
	var participants []ChannelParticipant

	if userIdList != nil && len(*userIdList) > 0  {
		err := db.Where("user_id IN ? AND channel_id = ?", *userIdList, channelId).
			Find(&participants).Error

		if err != nil {
			return nil, errors.New("error fetching a mentioned participant")
		}
	}

	return &participants, nil
}

func NewRepository() Repository {
	return &repositoryDefinition{}
}