package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
)

type BookUseCase interface {
	BookListUseCase(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	BookFindUseCase(id int64) (*[]model.Book, service.RecodeNotFoundError)
	DescriptionUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError)
}

type bookUseCase struct {
	BookRepo repository.BookRepository
}

func NewBookUseCase(cr repository.BookRepository) BookUseCase {
	return &bookUseCase{
		BookRepo: cr,
	}
}

func (u *bookUseCase) BookListUseCase(account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
	books, err := u.BookRepo.List(account)
	return books, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) BookFindUseCase(id int64) (*[]model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.Find(id)
	return book, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) DescriptionUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError) {
	book, err := u.BookRepo.Description(id)
	return book, service.RecodeNotFoundError(err)
}
