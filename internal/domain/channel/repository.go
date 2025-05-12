package channel

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindById(db *gorm.DB, id int64) (*Channel, error)
	FindOne(db *gorm.DB, where *OptionalChannelInterface) (*Channel, error)
	Find(db *gorm.DB, where *OptionalChannelInterface) (*[]Channel, error)
	Create(db *gorm.DB, payload *ChannelInterface) error
	Update(db *gorm.DB, channel *Channel) error

	CreateParticipant(userId string, channel *Channel) error
}

type repositoryDefinition struct{}

func (r *repositoryDefinition) FindById(db *gorm.DB, id int64) (*Channel, error) {
	var channel Channel

	if err := db.First(&channel, id).Error; err != nil {
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

	if err := db.Where(query).First(&channel).Error; err != nil {
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

	if err := db.Where(query).Find(&channels).Error; err != nil {
		return nil, err
	}

	return &channels, nil	
}

func (r *repositoryDefinition) Create(db *gorm.DB, payload *ChannelInterface) error {
	channel := Channel{Name: payload.Name, IsDirect: payload.IsDirect, Description: &payload.Description}

	result := db.Create(&channel).Error

	if result != nil {
		return errors.New("failed to create channel: " + result.Error())
	}

	return nil
}

func (r *repositoryDefinition) Update(db *gorm.DB, channel *Channel) error {
	result := db.Updates(channel).Error

	if result != nil {
		return errors.New("failed to update channel: " + result.Error())
	}

	return nil
}

func (r *repositoryDefinition) CreateParticipant(userId string, channel *Channel, participant ) error {
	result := db.Create()
}

func NewRepository() Repository {
	return &repositoryDefinition{}
}