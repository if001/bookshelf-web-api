package repository

import "bookshelf-web-api/domain/model"

type DescriptionRepository interface {
	FindDescriptions(id int64) (*[]model.Description, error)
	CreateDescription(description model.Description) (*model.Description, error)
	GetDescription(description model.Description) (*model.Description, error)
	Update(description model.Description) (*model.Description, error)
}


