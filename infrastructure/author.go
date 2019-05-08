package infrastructure

import (
	"bookshelf-web-api/domain/repository"
	"github.com/jinzhu/gorm"
)

type authorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &authorRepository{ DB : db }
}
