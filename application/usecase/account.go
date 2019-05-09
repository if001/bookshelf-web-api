package usecase


import (
	"bookshelf-web-api/domain/repository"
	"context"
	"bookshelf-web-api/domain/model"
)

type AccountUseCase interface {
	GetAccountUseCase(ctx context.Context) (*model.Account, error)
	// AccountGetUseCase(token string) (*tables.Account, error)
	SetAccountToCtxByToken(token string, ctx *context.Context) error
}

type accountUseCase struct {
	AccountRepo repository.AccountRepository
}

func NewAccountUseCase(cr repository.AccountRepository) AccountUseCase {
	return &accountUseCase {
		AccountRepo: cr,
	}
}

func (u *accountUseCase) GetAccountUseCase(ctx context.Context) (*model.Account, error) {
	account, err := u.AccountRepo.GetAccount(ctx)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (u *accountUseCase) SetAccountToCtxByToken(token string, ctx *context.Context) error {
	err := u.AccountRepo.SetAccount(token, ctx)
	if err != nil {
		return err
	}
	return nil
}