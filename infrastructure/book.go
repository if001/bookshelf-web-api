package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/infrastructure/tables"
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
		bookModel.Author = authorModel

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
//func createAuthor(r *bookRepository, tx *gorm.DB, author string) (*model.Author, error) {
//	var newAuthor model.Author
//	newAuthor.Name = author
//	err := tx.Create(&newAuthor).Error
//	if err != nil {
//		return nil, err
//	}
//	return &newAuthor, nil
//}
//
//func createCategories(r *bookRepository, tx *gorm.DB, categories []string) error {
//	var categoriesModel []model.Category
//
//	for i := range categories {
//		err := r.DB.Where("name = ?", categories[i]).Find(&categoriesModel).Error
//		if err != nil {
//			return err
//		}
//		if len(categoriesModel) == 0 {
//			newCategory := model.Category{}
//			newCategory.Name = categories[i]
//			err = tx.Create(&newCategory).Error
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
//
//func (r *bookRepository) CreateBook(bookRequest model.BookRequest, account model.Account) (*model.Book, service.RecodeNotFoundError) {
//	var authorModel []model.Author
//
//	err := r.DB.Where("name = ?",bookRequest.Author).Find(&authorModel).Error
//	if err != nil {
//		return nil, err
//	}
//
//	tx := r.DB.Begin()
//	defer func() {
//		err := recover()
//		if err != nil {
//			tx.Rollback()
//		}
//	}()
//
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
//	now := time.Now()
//	book := model.Book{}
//	book.AccountID = account.ID
//	book.Title = bookRequest.Title
//	book.AuthorID = authorId
//	book.PublishedAt = mysql.NullTime{Time:now, Valid:false }
//	book.StartAt = mysql.NullTime{Time:now, Valid:false }
//	book.EndAt = mysql.NullTime{Time:now, Valid:false }
//
//	err = tx.Create(&book).Error
//	if err != nil {
//		tx.Rollback()
//		return nil, err
//	}
//	err = tx.Commit().Error
//	return &book, err
//}
//
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
