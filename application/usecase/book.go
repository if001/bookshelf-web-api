package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
)

type BookUseCase interface {
	BookListUseCase(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	BookFindUseCase(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	DescriptionUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError)
	CreateBook(bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError)
	UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError)
}

type bookUseCase struct {
	BookRepo repository.BookRepository
	AuthorRepo repository.AuthorRepository
}

func NewBookUseCase(br repository.BookRepository, ar repository.AuthorRepository) BookUseCase {
	return &bookUseCase{
		BookRepo: br,
		AuthorRepo: ar,
	}
}

func (u *bookUseCase) BookListUseCase(account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
	books, err := u.BookRepo.List(account)
	return books, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) BookFindUseCase(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.Find(id, account)
	return book, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) DescriptionUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError) {
	book, err := u.BookRepo.Description(id)
	return book, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) CreateBook(bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.Create(bookRequest, account)
	return book, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.Update(id, bookRequest, account)
	return book, service.RecodeNotFoundError(err)
}
