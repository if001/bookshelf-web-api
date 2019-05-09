package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type CategoryRepository interface {
	Get() (*model.Category, service.RecodeNotFoundError)
	GetByBookId(bookId int64) (*[]model.Category, error)
	GetNotExistCategories(categories []model.Category) (*[]model.Category, error)
}

