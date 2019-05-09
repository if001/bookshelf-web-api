package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/infrastructure/tables"
	"database/sql"
	"github.com/go-sql-driver/mysql"
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

func (r *bookRepository) CreateBook(book model.Book, account model.Account) (*model.Book, error) {
	tx := r.DB.Begin()
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
		}
	}()

	var authorId = sql.NullInt64{ Valid:false }
	if book.Author != nil {
		var authorTable = tables.Author{}
		authorTable.Name = book.Author.Name
		newAuthor, err := createAuthor(tx, authorTable)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		authorId = sql.NullInt64{Int64:newAuthor.ID }
	}
	categoriesTable := []tables.Category{}
	for i := range book.Categories {
		categoryTable := tables.Category{}
		categoryTable.Name = book.Categories[i].Name
		categoriesTable = append(
			categoriesTable,
			categoryTable,
		)
	}

	err := createCategories(tx, categoriesTable)
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




//func (r *bookRepository) FindBook(id int64, account model.Account) (*[]model.Book, service.RecodeNotFoundError) {
//	var books []model.Book
//	var categoriesModel []model.Category
//	var descriptions []model.Description
//
//	err := r.DB.Where("account_id = ?", account.ID).Where("id = ?",id).Find(&books).Error
//	if err != nil {
//		return &[]model.Book{}, err
//	}
//	for i := range books {
//		if books[i].AuthorID.Int64 != 0 {
//			err = r.DB.Model(books[i]).Related(&books[i].Author,"Author").Error
//			if err != nil {
//				return &[]model.Book{}, err
//			}
//		} else {
//			books[i].Author = model.Author{}
//		}
//		err = r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
//			Where("book_id = ?", books[i].ID).
//			Find(&categoriesModel).
//			Error
//		if err != nil {
//			return &[]model.Book{}, err
//		}
//		books[i].Categories = categoriesModel
//
//		err = r.DB.Where("book_id = ?",books[i].ID).Find(&descriptions).Error
//		if err != nil {
//			return &[]model.Book{}, err
//		}
//		books[i].Description = descriptions
//	}
//	return &books, err
//}
//
//func (c *bookRepository) FindDescriptions(id int64) (*[]model.Description, service.RecodeNotFoundError) {
//	var descriptions = []model.Description{}
//
//	err := c.DB.Where("book_id = ?", id).Find(&descriptions).Error
//	return &descriptions, err
//}
//
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
func createAuthor(tx *gorm.DB, authorTable tables.Author) (*model.Author, error) {
	err := tx.Create(&authorTable).Error
	if err != nil {
		return nil, err
	}
	newAuthor := model.Author{}
	newAuthor.ID = authorTable.ID
	newAuthor.Name = authorTable.Name
	newAuthor.CreatedAt = authorTable.CreatedAt
	newAuthor.UpdatedAt = authorTable.UpdatedAt
	return &newAuthor, nil
}

func createCategories(tx *gorm.DB, categories []tables.Category) error {
	for i := range categories {
		err := tx.Create(&categories[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}


//func (r *bookRepository) UpdateBook(id int64, bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
//	var authorModel []model.Author
//
//	err := r.DB.Where("name = ?",bookRequest.Author).Find(&authorModel).Error
//	if err != nil {
//		return &model.Book{}, err
//	}
//
//	tx := r.DB.Begin()
//	defer func() {
//		err := recover()
//		if err != nil {
//			tx.Rollback()
//		}
//	}()
//	var authorId = sql.NullInt64{Int64:0, Valid:true }
//	if len(authorModel) == 0 {
//		newAuthor, err := createAuthor(r, tx, bookRequest.Author)
//		if err != nil {
//			tx.Rollback()
//			return nil, err
//		}
//		authorId = sql.NullInt64{Int64:newAuthor.ID, Valid:true }
//	} else {
//		authorId = sql.NullInt64{Int64:authorModel[0].ID, Valid:true }
//	}
//
//	err = createCategories(r, tx, bookRequest.Categories)
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//
//	books := []model.Book{}
//	err = r.DB.Where("id = ?", id).Find(&books).Error
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//	if len(books) == 0 {
//		tx.Rollback()
//		return nil, errors.New("record not found")
//	}
//	book := books[0]
//
//	if bookRequest.Title != "" {
//		book.Title = bookRequest.Title
//	}
//	book.AuthorID = authorId
//	if bookRequest.PrevBookId != 0 {
//		book.PrevBookID.Int64 = bookRequest.PrevBookId
//	}
//	if bookRequest.NextBookId != 0 {
//		book.NextBookID.Int64 = bookRequest.NextBookId
//	}
//	err = tx.Create(&book).Error
//	if err != nil {
//			tx.Rollback()
//		return nil, err
//	}
//
//	err = tx.Commit().Error
//	return &book, err
//}
