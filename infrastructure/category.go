package infrastructure

import (
	"bookshelf-web-api/domain/repository"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{ DB : db }
}

var category model.Category

func (c *categoryRepository) Get() (*model.Category, service.RecodeNotFoundError) {
	var err  = c.DB.Find(&category).Error
	return &category, err
}