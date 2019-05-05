package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type AccountRepository interface {
	Get(token string) (*model.Account, service.RecodeNotFoundError)
}
