package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type DescriptionRepository interface {
	Update(id int64, description model.DescriptionRequest) (*model.Description, service.RecodeNotFoundError)
}


