package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/infrastructure/tables"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"errors"
)

// TODO DBしかうけとってないしリポジトリまとめたほうがよい？
type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{ DB : db }
}

func (r *bookRepository) GetBooks(accountId int64) (*[]model.Book, error) {
	var bookTables []tables.Book
	var categoriesTable []tables.Category
	var descriptionTables []tables.Description
	var bookModels []model.Book

	err := r.DB.Where("account_id = ?", accountId).Find(&bookTables).Error
	if err != nil {
		return nil, err
	}
	for i := range bookTables {
		var bookModel = model.Book{}
		var authorModel = model.Author{}
		var categoryModel = model.Category{}
		var descriptionModel = model.Description{}

		if bookTables[i].AuthorID.Int64 != 0 {
			err = r.DB.Model(bookTables[i]).Related(&bookTables[i].Author,"Author").Error
			if err != nil {
				return nil, err
			}
			authorModel.Fill(
				bookTables[i].Author.ID,
				bookTables[i].Author.Name,
				bookTables[i].Author.CreatedAt,
				bookTables[i].Author.UpdatedAt,
			)
		}
		bookModel.Author = &authorModel

		err = r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
			Where("book_id = ?", bookTables[i].ID).
			Find(&categoriesTable).
			Error
		if err != nil {
			return nil, err
		}
		for i := range categoriesTable {
			categoryModel.Fill(
				categoriesTable[i].ID,
				categoriesTable[i].Name,
				categoriesTable[i].CreatedAt,
				categoriesTable[i].UpdatedAt,
			)
			bookModel.Categories = append(
				bookModel.Categories,
				categoryModel)
		}

		err = r.DB.Where("book_id = ?", bookTables[i].ID).Find(&descriptionTables).Error
		if err != nil {
			return nil, err
		}
		bookTables[i].Description = descriptionTables
		for i := range descriptionTables {
			descriptionModel.Fill(
				descriptionTables[i].ID,
				descriptionTables[i].Description,
				descriptionTables[i].CreatedAt,
				descriptionTables[i].UpdatedAt,
			)
			bookModel.Descriptions = append(
				bookModel.Descriptions,
				descriptionModel)
		}
		bookModels = append(bookModels, bookModel)
	}
	return &bookModels, err
}

func isIncludeCategory(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func updateBookCategory(db *gorm.DB,tx *gorm.DB, book model.Book) (error) {
	ptrue := &[]bool{true}[0]
	pfalse := &[]bool{false}[0]

	existBookCategoriesTable := []tables.BookCategory{}
	err := db.Where("book_id = ?", book.ID).Find(&existBookCategoriesTable).Error
	if err != nil {
		return err
	}
	existIds := []int64{}
	for i := range existBookCategoriesTable {
		existIds = append(
			existIds,
			existBookCategoriesTable[i].ID,
		)
	}

	requestCategories := book.Categories
	requestIds := []int64{}
	for i := range requestCategories {
		requestIds = append(
			requestIds,
			requestCategories[i].ID,
		)
	}

	for i := range existBookCategoriesTable {
		if !isIncludeCategory(existBookCategoriesTable[i].ID, requestIds) {
			bookCategory := tables.BookCategory{}
			bookCategory.ID = existBookCategoriesTable[i].ID
			bookCategory.BookID = book.ID
			bookCategory.Status = pfalse
			err := tx.Save(&bookCategory).Error
			if err != nil {
				return err
			}
		}
	}

	for i := range requestCategories {
		if isIncludeCategory(requestCategories[i].ID, existIds) {
			bookCategory := tables.BookCategory{}
			bookCategory.ID = requestCategories[i].ID
			bookCategory.BookID = book.ID
			bookCategory.Status = ptrue
			err := tx.Create(&bookCategory).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *bookRepository) CreateBook(book model.Book, account model.Account) (result *model.Book, err error) {
	tx := r.DB.Begin()
	defer func() {
		rcv := recover()
		if rcv != nil {
			err = tx.Rollback().Error
			if err == nil {
				err = errors.New("in recover: "+rcv.(string))
			}
		}
	}()

	authorId := sql.NullInt64{Int64:book.Author.ID, Valid: book.Author != nil }

	err = updateBookCategory(r.DB,tx, book)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	bookTable := tables.Book{}
	bookTable.AccountID = account.ID
	bookTable.Title = book.Name
	bookTable.AuthorID = authorId
	bookTable.PublishedAt = mysql.NullTime{Valid:false }
	bookTable.StartAt = mysql.NullTime{Valid:false }
	bookTable.EndAt = mysql.NullTime{Valid:false }
	bookTable.NextBookID = sql.NullInt64{Int64:book.NextBookID, Valid:book.NextBookID != 0}
	bookTable.PrevBookID = sql.NullInt64{Int64:book.PrevBookID, Valid:book.PrevBookID != 0}

	err = tx.Create(&bookTable).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &book, nil
}


func (r *bookRepository) FindBook(id int64, account model.Account) (*model.Book, error) {
	bookTable := []tables.Book{}
	authorTable := []tables.Author{}
	categoriesTable := []tables.Category{}
	descriptionsTable := []tables.Description{}
	authorModel := model.Author{}
	categoriesModel := []model.Category{}
	descriptionsModel := []model.Description{}

	err := r.DB.Where("account_id = ?", account.ID).Where("id = ?",id).Find(&bookTable).Error
	if err != nil {
		return nil, err
	}
	if len(bookTable) == 0 {
		return nil, errors.New("table not found")
	}

	err = r.DB.Joins("JOIN books ON books.author_id = author.id").
		Where("books.id = ?", bookTable[0].ID).
		Find(&authorTable).
		Error
	if err != nil {
		return nil, err
	}
	if len(authorTable) != 0 {
		authorModel.Fill(
			authorTable[0].ID,
			authorTable[0].Name,
			authorTable[0].CreatedAt,
			authorTable[0].UpdatedAt,
		)
	} else {
		authorModel = model.Author{}
	}

	err = r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
		Where("book_id = ?", bookTable[0].ID).
		Find(&categoriesTable).
		Error
	if err != nil {
		return nil, err
	}

	for i := range categoriesTable{
		category := model.Category{}
		category.Fill(
			categoriesTable[i].ID,
			categoriesTable[i].Name,
			categoriesTable[i].CreatedAt,
			categoriesTable[i].UpdatedAt,
		)
		categoriesModel = append(
			categoriesModel,
			category,
		)
	}
	err = r.DB.Where("book_id = ?",bookTable[0].ID).Find(&descriptionsTable).Error
	if err != nil {
		return nil, err
	}
	bookTable[0].Description = descriptionsTable
	for i := range descriptionsTable{
		description := model.Description{}
		description.Fill(
			descriptionsTable[i].ID,
			descriptionsTable[i].Description,
			descriptionsTable[i].CreatedAt,
			descriptionsTable[i].UpdatedAt,
		)
		descriptionsModel = append(
			descriptionsModel,
			description,
		)
	}
	book := model.Book{}

	book.Fill(
		bookTable[0].ID,
		bookTable[0].Title,
		&authorModel,
		bookTable[0].PublishedAt,
		"publisher", //TODO あとでテーブルにからむつくる
		account.ID,
		bookTable[0].StartAt,
		bookTable[0].EndAt,
		bookTable[0].NextBookID.Int64,
		bookTable[0].PrevBookID.Int64,
		descriptionsModel,
		categoriesModel,
		bookTable[0].CreatedAt,
		bookTable[0].UpdatedAt,
	)
	return &book, err
}


//func (c *bookRepository) CreateDescription(id int64, description string) (*model.Description, service.RecodeNotFoundError) {
//	var books []model.Book
//
//	err := c.DB.Where("id = ?", id).Find(&books).Error
//	if err != nil {
//		return nil, err
//	}
//	if len(books) == 0 {
//		return nil, errors.New("record not found")
//	}
//
//	newDescription := model.Description{}
//	newDescription.BookId = books[0].ID
//	newDescription.Description = description
//	err = c.DB.Create(&newDescription).Error
//	if err != nil {
//		return nil, err
//	}
//	return &newDescription, err
//}
//
//


func (r *bookRepository) UpdateBook(book model.Book, account model.Account) (result *model.Book, err error) {
	result = nil
	err = nil

	tx := r.DB.Begin()
	defer func() {
		rcv := recover()
		if rcv != nil {
			err = tx.Rollback().Error
			if err == nil {
				err = errors.New("in recover: "+rcv.(string))
			}
		}
	}()

	bookTable := []tables.Book{}
	err = r.DB.Where("id = ?", book.ID).Find(&bookTable).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if len(bookTable) == 0 {
		tx.Rollback()
		err = errors.New("record not found")
		return
	}

	authorId := sql.NullInt64{Int64:book.Author.ID, Valid: book.Author != nil }

	err = updateBookCategory(r.DB,tx, book)
	if err != nil {
		tx.Rollback()
		return
	}

	if book.Name != "" {
		bookTable[0].Title = book.Name
	}
	bookTable[0].AuthorID = authorId
	bookTable[0].StartAt = mysql.NullTime{Valid:false }
	bookTable[0].EndAt = mysql.NullTime{Valid:false }
	bookTable[0].NextBookID = sql.NullInt64{Int64:book.NextBookID, Valid:book.NextBookID != 0}
	bookTable[0].PrevBookID = sql.NullInt64{Int64:book.PrevBookID, Valid:book.PrevBookID != 0}

	err = tx.Save(&bookTable[0]).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}
	result = &book
	return
}
