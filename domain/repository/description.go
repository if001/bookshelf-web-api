package repository

import "bookshelf-web-api/domain/model"

type DescriptionRepository interface {
	FindDescriptions(id int64) (*[]model.Description, error)
	// Update(id int64, description model.DescriptionRequest) (*model.Description, service.RecodeNotFoundError)
}


