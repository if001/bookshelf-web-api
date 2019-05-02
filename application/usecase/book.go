package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
)

type BookUseCase interface {
	BookListUseCase() (*[]model.Book, error)
}

type bookUseCase struct {
	BookRepo repository.BookRepository
}

func NewBookUseCase(cr repository.BookRepository) BookUseCase {
	return &bookUseCase{
		BookRepo: cr,
	}
}

func (u *bookUseCase) BookListUseCase() (*[]model.Book, error) {
	books, err := u.BookRepo.List()
	return books, err
}
	return cate, err
}
