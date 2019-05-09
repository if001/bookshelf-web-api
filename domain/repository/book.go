package repository

import "bookshelf-web-api/domain/model"

type BookRepository interface {
	GetBooks(accountId int64) (* []model.Book, error)
	CreateBook(bookRequest model.Book,account model.Account) (*model.Book, error)

	FindBook(id int64, account model.Account) (*model.Book, error)
	//GetBooks(account model.Account) (*[]model.Book, service.RecodeNotFoundError)

	//UpdateBook(id int64, bookRequest model.BookRequest,account model.Account) (*model.Book, service.RecodeNotFoundError)
	//FindDescriptions(id int64) (*[]model.Description, service.RecodeNotFoundError)
	//CreateDescription(id int64, description string) (*model.Description, service.RecodeNotFoundError)
}
