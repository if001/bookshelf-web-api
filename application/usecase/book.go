package usecase

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
)

type BookUseCase interface {
	BookListUseCase(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	BookFindUseCase(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	DescriptionFindUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError)
	DescriptionCreateUseCase(id int64, description string) (*model.Description, service.RecodeNotFoundError)
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
	books, err := u.BookRepo.GetBooks(account)
	return books, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) BookFindUseCase(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.FindBook(id, account)
	return book, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) DescriptionFindUseCase(id int64) (*[]model.Description, service.RecodeNotFoundError) {
	book, err := u.BookRepo.FindDescriptions(id)
	return book, service.RecodeNotFoundError(err)
}

func (u *bookUseCase) DescriptionCreateUseCase(id int64, description string) (*model.Description, service.RecodeNotFoundError) {
	newDescription, err := u.BookRepo.CreateDescription(id, description)
	return newDescription, service.RecodeNotFoundError(err)
}

func (u *bookUseCase) CreateBook(bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.CreateBook(bookRequest, account)
	return book, service.RecodeNotFoundError(err)
}
func (u *bookUseCase) UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
	book, err := u.BookRepo.UpdateBook(id, bookRequest, account)
	return book, service.RecodeNotFoundError(err)
}
