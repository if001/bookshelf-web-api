package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// TODO DBしかうけとってないしリポジトリまとめたほうがよい？
type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{ DB : db }
}

var book model.Book
var books []model.Book
var categoriesModel []model.Category
var descriptions []model.Description

func (r *bookRepository) List(account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
	err := r.DB.Where("account_id = ?", account.ID).Find(&books).Error
	if err != nil {
		return &[]model.Book{}, err 
	}
	for i := range books {
		if books[i].AuthorID.Int64 != 0 {
			err = r.DB.Model(books[i]).Related(&books[i].Author,"Author").Error
			if err != nil {
				return &[]model.Book{}, err 
			}
		} else {
			books[i].Author = model.Author{}
		}
		err = r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
			Where("book_id = ?", books[i].ID).
			Find(&categoriesModel).
			Error
		if err != nil {
			return &[]model.Book{}, err 
		}
		books[i].Categories = categoriesModel

		err = r.DB.Where("book_id = ?",books[i].ID).Find(&descriptions).Error
		if err != nil {
			return &[]model.Book{}, err 
		}
		books[i].Description = descriptions
	}
	return &books, err
}

func (r *bookRepository) Find(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
	err := r.DB.Where("account_id = ?", account.ID).Where("id = ?",id).Find(&books).Error
	if err != nil {
		return &[]model.Book{}, err 
	}
		for i := range books {
		if books[i].AuthorID.Int64 != 0 {
			err = r.DB.Model(books[i]).Related(&books[i].Author,"Author").Error
			if err != nil {
				return &[]model.Book{}, err 
			}
		} else {
			books[i].Author = model.Author{}
		}
		err = r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
			Where("book_id = ?", books[i].ID).
			Find(&categoriesModel).
			Error
		if err != nil {
			return &[]model.Book{}, err 
		}
		books[i].Categories = categoriesModel

		err = r.DB.Where("book_id = ?",books[i].ID).Find(&descriptions).Error
		if err != nil {
			return &[]model.Book{}, err 
		}
		books[i].Description = descriptions
	}
	return &books, err
}

func (c *bookRepository) Description(id int64) (*[]model.Description, service.RecodeNotFoundError) {
	err := c.DB.Where("book_id = ?", id).Find(&descriptions).Error
	return &descriptions, err
}

func createAuthor(r *bookRepository, tx *gorm.DB, author string) (*model.Author, error) {
	newAuthor := model.Author{}
	newAuthor.Name = author
	err := tx.Create(&newAuthor).Error
	if err != nil {
		return nil, err
	}
	return &newAuthor, nil
}

func createCategories(r *bookRepository, tx *gorm.DB, categories []string) error {
	for i := range categories {
		err := r.DB.Where("name = ?", categories[i]).Find(&categoriesModel).Error
		if err != nil {
			return err
		}
		if len(categoriesModel) == 0 {
			newCategory := model.Category{}
			newCategory.Name = categories[i]
			err = tx.Create(&newCategory).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

var authorModel []model.Author
func (r *bookRepository) Create(bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
	err := r.DB.Where("name = ?",bookRequest.Author).Find(&authorModel).Error
	if err != nil {
		return &model.Book{}, err
	}

	tx := r.DB.Begin()
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		}
	}()

	var authorId = sql.NullInt64{Int64:0, Valid:true }
	if len(authorModel) == 0 {
		newAuthor, err := createAuthor(r, tx, bookRequest.Author)
		if err != nil {
			tx.Rollback()
			return &model.Book{}, err
		}
		authorId = sql.NullInt64{Int64:newAuthor.ID, Valid:true }
	} else {
		authorId = sql.NullInt64{Int64:authorModel[0].ID, Valid:true }
	}

	err = createCategories(r, tx, bookRequest.Categories)
	if err != nil {
		tx.Rollback()
		return &model.Book{}, err
	}
	
	now := time.Now()
	book := model.Book{}
	book.AccountID = account.ID
	book.Title = bookRequest.Title
	book.AuthorID = authorId
	book.PublishedAt = mysql.NullTime{Time:now, Valid:false }
	book.StartAt = mysql.NullTime{Time:now, Valid:false }
	book.EndAt = mysql.NullTime{Time:now, Valid:false }

	err = tx.Create(&book).Error
	if err != nil {
		tx.Rollback()
		return &model.Book{}, err
	}
	err = tx.Commit().Error
	return &book, err
}

func (r *bookRepository) Update(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
	err := r.DB.Where("name = ?",bookRequest.Author).Find(&authorModel).Error
	if err != nil {
		return &model.Book{}, err
	}

	tx := r.DB.Begin()
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		}
	}()
	var authorId = sql.NullInt64{Int64:0, Valid:true }
	if len(authorModel) == 0 {
		newAuthor, err := createAuthor(r, tx, bookRequest.Author)
		if err != nil {
			tx.Rollback()
			return &model.Book{}, err
		}
		authorId = sql.NullInt64{Int64:newAuthor.ID, Valid:true }
	} else {
		authorId = sql.NullInt64{Int64:authorModel[0].ID, Valid:true }
	}

	err = createCategories(r, tx, bookRequest.Categories)
	if err != nil {
		tx.Rollback()
		return &model.Book{}, err
	}

	err = r.DB.Where("id = ?", id).Find(&book).Error
	if err != nil {
		tx.Rollback()
		return &model.Book{}, err
	}


	if bookRequest.Title != "" {
		book.Title = bookRequest.Title
	}
	book.AuthorID = authorId
	if bookRequest.PrevBookId != 0 {
		book.PrevBookID.Int64 = bookRequest.PrevBookId
	}
	if bookRequest.NextBookId != 0 {
		book.NextBookID.Int64 = bookRequest.NextBookId
	}
	err = tx.Create(&book).Error
	if err != nil {
			tx.Rollback()
		return &model.Book{}, err
	}

	err = tx.Commit().Error
	return &book, err
}
