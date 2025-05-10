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
}

type repository struct{}

func (r *repository) FindById(db *gorm.DB, id int64) (*Channel, error) {
	var channel Channel

	if err := db.First(&channel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Channel not found")
		}
		return nil, err
	}

	return &channel, nil
}

func (r *repository) FindOne(db *gorm.DB, where *OptionalChannelInterface) (*Channel, error) {
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

func (r *repository) Find(db *gorm.DB, where *OptionalChannelInterface) (*[]Channel, error) {
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

func (r *repository) Create(db *gorm.DB, payload *ChannelInterface) error {
	channel := Channel{Name: payload.Name, IsDirect: payload.IsDirect, Description: &payload.Description}

	result := db.Create(&channel).Error

	if result != nil {
		return errors.New("failed to create channel: " + result.Error())
	}

	return nil
}

func NewRepository() Repository {
	return &repository{}
}