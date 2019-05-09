package repository

import "bookshelf-web-api/domain/model"

type AuthorRepository interface {
	IsExistAuthor(author *model.Author) (bool, error)
	GetAuthor(id int64) (*model.Author, error)
}