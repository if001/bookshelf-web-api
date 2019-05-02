package repository

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
)

type BookRepository interface {
	List() (*model.Book, service.RecodeNotFoundError)
}


