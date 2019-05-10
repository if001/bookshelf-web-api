package repository

import "bookshelf-web-api/domain/model"

type BookRepository interface {
	GetBooks(accountId int64) (* []model.Book, error)
	CreateBook(bookRequest model.Book,account model.Account) (*model.Book, error)

	FindBook(id int64, account model.Account) (*model.Book, error)
	UpdateBook(book model.Book, account model.Account) (*model.Book, error)

	//GetBooks(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
}
