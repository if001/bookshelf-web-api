package usecase


import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
)

type AccountUseCase interface {
	AccountGetUseCase(token string) (*model.Account, service.RecodeNotFoundError)
}

type accountUseCase struct {
	AccountRepo repository.AccountRepository
}

func NewAccountUseCase(cr repository.AccountRepository) AccountUseCase {
	return &accountUseCase {
		AccountRepo: cr,
	}
}

func (u *accountUseCase) AccountGetUseCase(token string) (*model.Account, service.RecodeNotFoundError) {
	account, err := u.AccountRepo.Get(token)
	return account, service.RecodeNotFoundError(err)
}
