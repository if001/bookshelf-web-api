package repository

import (
	"bookshelf-web-api/infrastructure/tables"
	"context"
)

type AccountRepository interface {
	GetAccount(ctx context.Context) (*tables.Account, error)
	SetAccount(token string, ctx *context.Context) (error)
}
