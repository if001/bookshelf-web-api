package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
)

type BookUseCase interface {
	BookListUseCase() (*[]model.Book, service.RecodeNotFoundError)
	BookFindUseCase(id int64) (*[]model.Book, service.RecodeNotFoundError)
}

type bookUseCase struct {
	BookRepo repository.BookRepository
}

func NewBookUseCase(cr repository.BookRepository) BookUseCase {
	return &bookUseCase{
		BookRepo: cr,
	}
}

func (u *bookUseCase) BookListUseCase() (*[]model.Book, service.RecodeNotFoundError) {
	books, err := u.BookRepo.List()
	return books, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) BookFindUseCase(id int64) (*[]model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.Find(id)
	return book, service.RecodeNotFoundError(err)
}
	return cate, err
}
