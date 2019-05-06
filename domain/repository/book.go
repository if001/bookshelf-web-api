package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type BookRepository interface {
	List(account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	Find(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError)
	Description(id int64) (*[]model.Description, service.RecodeNotFoundError)
	Create(bookRequest model.BookRequest,account model.Account) (*model.Book, service.RecodeNotFoundError)
}