package repository

import "bookshelf-web-api/domain/model"

type BookRepository interface {
	GetBooks(account model.Account) (* []model.Book, error)
	CreateBook(bookRequest model.Book,account model.Account) (*model.Book, error)

	FindBook(id int64, account model.Account) (*model.Book, error)
	UpdateBook(book model.Book, account model.Account) (*model.Book, error)

	StartReadBook(book model.Book) (*model.Book, error)
	EndReadBook(book model.Book) (*model.Book, error)
	//GetBooks(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
}
