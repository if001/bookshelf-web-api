package repository

import (
	"bookshelf-web-api/domain/model"
)

type CategoryRepository interface {
	GetCategories() (*[]model.Category, error)
	GetByIds(categoryIds []int64) (*[]model.Category, error)
	GetByBookId(bookId int64) (*[]model.Category, error)
	GetNotExistCategories(categories []model.Category) (*[]model.Category, error)
	LogicalDelete(bookId int64, categoryId int64)  (error)
}

