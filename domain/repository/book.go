package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type BookRepository interface {
	GetBooks(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	FindBook(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	CreateBook(bookRequest model.BookRequest,account model.Account) (*model.Book, service.RecodeNotFoundError)
	UpdateBook(id int64, bookRequest model.BookRequest,account model.Account) (*model.Book, service.RecodeNotFoundError)
	GetDescriptions(id int64) (*[]model.Description, service.RecodeNotFoundError)
}
