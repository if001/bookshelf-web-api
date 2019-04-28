package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type CategoryRepository interface {
	Get() (*model.Category, service.RecodeNotFoundError)
}

