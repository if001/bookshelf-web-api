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
var books []model.Book

func (c *bookRepository) List() (*[]model.Book, service.RecodeNotFoundError) {
	err := c.DB.Find(&books).Error
	return &books, err
}

func (c *bookRepository) Find(id int64) (*[]model.Book, service.RecodeNotFoundError) {
	err := c.DB.Where("id = ?",id).Find(&books).Error
	return &books, err
}

var descriptions []model.Description
func (c *bookRepository) Description(id int64) (*[]model.Description, service.RecodeNotFoundError) {
	err := c.DB.Where("book_id = ?", id).Find(&descriptions).Error
	return &descriptions, err
}