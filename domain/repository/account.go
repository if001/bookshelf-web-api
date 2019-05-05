package repository

import (
	"bookshelf-web-api/domain/model"
)

type AccountRepository interface {
	Get(token string) (*model.Account, error)
}
