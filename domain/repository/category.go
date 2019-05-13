package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type CategoryRepository interface {
	GetCategories() (*[]model.Category, service.RecodeNotFoundError)
	GetByIds(categoryIds []int64) (*[]model.Category, service.RecodeNotFoundError)
	GetByBookId(bookId int64) (*[]model.Category, error)
	GetNotExistCategories(categories []model.Category) (*[]model.Category, error)
}

