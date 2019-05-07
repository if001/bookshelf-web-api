package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"github.com/jinzhu/gorm"
)

type descriptionRepository struct {
	DB *gorm.DB
}

func NewDescriptionRepository(db *gorm.DB) repository.DescriptionRepository {
	return &descriptionRepository{ DB : db }
}

var descriptionModel model.Description


func (c *descriptionRepository) Update(id int64, description string) (*model.Description, service.RecodeNotFoundError) {
	err := c.DB.Where("id = ?", id).Find(&descriptionModel).Error
	if err != nil {
		return nil, err
	}
	descriptionModel.Description = description
	err = c.DB.Save(&descriptionModel).Error
	if err != nil {
		return nil, err
	}
	return &descriptionModel, err
}
