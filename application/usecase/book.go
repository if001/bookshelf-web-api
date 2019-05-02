package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
)

type BookUseCase interface {
	BookUseCase() (*model.Book, error)
}

type bookUseCase struct {
	BookRepo repository.BookRepository
}

func NewBookUseCase(cr repository.BookRepository) BookUseCase {
	return &bookUseCase{
		BookRepo: cr,
	}
}

func (u *bookUseCase) BookUseCase() (*model.Book, error) {
	cate, err := u.BookRepo.List()
	return cate, err
}
