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

var books []model.Book
var categories []model.Category
var descriptions []model.Description

func (c *bookRepository) List() (*[]model.Book, service.RecodeNotFoundError) {
	accountId := 1
	err := c.DB.Where("account_id = ?", accountId).Find(&books).Error
	for i := range books {
		if books[i].AuthorID != 0 {
			err = c.DB.Model(books[i]).Related(&books[i].Author,"Author").Error
		} else {
			books[i].Author = model.Author{}
		}
		err = c.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
			Where("book_id = ?", books[i].ID).
			Find(&categories).
			Error
		books[i].Categories = categories

		err = c.DB.Where("book_id = ?",books[i].ID).Find(&descriptions).Error
		books[i].Description = descriptions
	}
	return &books, err
}

func (c *bookRepository) Find(id int64) (*[]model.Book, service.RecodeNotFoundError) {
	err := c.DB.Where("id = ?",id).Find(&books).Error
	return &books, err
}

func (c *bookRepository) Description(id int64) (*[]model.Description, service.RecodeNotFoundError) {
	err := c.DB.Where("book_id = ?", id).Find(&descriptions).Error
	return &descriptions, err
}