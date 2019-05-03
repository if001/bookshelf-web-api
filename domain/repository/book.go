package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type BookRepository interface {
	List() (*[]model.Book, service.RecodeNotFoundError)
	Find(id int64) (*[]model.Book, service.RecodeNotFoundError)
	Description(id int64) (*[]model.Description, service.RecodeNotFoundError)
}


