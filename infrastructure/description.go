package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"errors"
	"github.com/jinzhu/gorm"
)

type descriptionRepository struct {
	DB *gorm.DB
}

func NewDescriptionRepository(db *gorm.DB) repository.DescriptionRepository {
	return &descriptionRepository{ DB : db }
}


func (c *descriptionRepository) Update(id int64, description string) (*[]model.Description, service.RecodeNotFoundError) {
	var descriptionModel []model.Description
	err := c.DB.Find(&descriptionModel,"id = ?", id).Error
	// err := c.DB.Where("id = ?", id).Find(&descriptionModel).Error
	if err != nil {
		return nil, err
	}
	if len(descriptionModel) == 0 {
		return nil, errors.New("not found")
	}
	descriptionModel[0].Description = description
	err = c.DB.Save(&descriptionModel).Error
	if err != nil {
		return nil, err
	}
	return &descriptionModel, err
}
