package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"github.com/jinzhu/gorm"
)

type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{ DB : db }
}

var book model.Book

func (c *bookRepository) List() (*model.Book, service.RecodeNotFoundError) {
	var err  = c.DB.Find(&book).Error
	return &book, err
}