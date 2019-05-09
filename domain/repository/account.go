package repository

import (
	"context"
	"bookshelf-web-api/domain/model"
)

type AccountRepository interface {
	GetAccount(ctx context.Context) (*model.Account, error)
	SetAccount(token string, ctx *context.Context) (error)
}
