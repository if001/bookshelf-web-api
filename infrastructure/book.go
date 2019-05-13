package infrastructure

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/repository"
	"bookshelf-web-api/domain/service"
	"github.com/jinzhu/gorm"
	"bookshelf-web-api/infrastructure/tables"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"errors"
	"time"
)

// TODO DBしかうけとってないしリポジトリまとめたほうがよい？
type bookRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{ DB : db }
}


//func (r *bookRepository) GetBooks(accountId int64) (*[]model.Book, error) {
//	var bookTablesAll []tables.Book2
//	err2 := r.DB.
//		Table("books").
//		Select("books.*, author.name as author_name, categories.*,description.description, description.created_at as descriptionCreatedAt").
//		Where("books.account_id = ?", accountId).
//		Where("books.id = ?", 2).
//		Joins("LEFT JOIN author ON books.author_id = author.id").
//		Joins("LEFT JOIN books_categories ON books_categories.book_id = books.id").
//		Joins("LEFT JOIN categories ON books_categories.category_id = categories.id").
//		Joins("LEFT JOIN description ON description.book_id = books.id").
//		Find(&bookTablesAll).
//		Error
//}


func (r *bookRepository) GetBooks(account model.Account) (*[]model.Book, error) {
	bookTable := []tables.Book{}
	books := []model.Book{}

	err := r.DB.Where("account_id = ?", account.ID).Find(&bookTable).Error
	if err != nil {
		return nil, err
	}
	if len(bookTable) == 0 {
		return nil, errors.New("table not found")
	}

	for i := range bookTable {
		authorTable := []tables.Author{}
		categoriesTable := []tables.Category{}
		descriptionsTable := []tables.Description{}

		authorModel := &model.Author{}
		categoriesModel := []model.Category{}
		descriptionsModel := []model.Description{}


		err = r.DB.Joins("JOIN books ON books.author_id = author.id").
			Where("books.id = ?", bookTable[i].ID).
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
			authorModel = nil
		}

		err = r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
			Where("book_id = ?", bookTable[i].ID).
			Where("status = ?", pTrue).
			Find(&categoriesTable).
			Error
		if err != nil {
			return nil, err
		}

		for i := range categoriesTable {
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
		err = r.DB.Where("book_id = ?", bookTable[i].ID).Find(&descriptionsTable).Error
		if err != nil {
			return nil, err
		}
		if len(descriptionsTable) != 0 {
			for i := range descriptionsTable {
				description := model.Description{}
				description.Fill(
					descriptionsTable[i].ID,
					descriptionsTable[i].BookId,
					descriptionsTable[i].Description,
					descriptionsTable[i].CreatedAt,
					descriptionsTable[i].UpdatedAt,
				)
				descriptionsModel = append(
					descriptionsModel,
					description,
				)
			}
		}

		book := model.Book{}

		book.Fill(
			bookTable[i].ID,
			bookTable[i].Title,
			authorModel,
			bookTable[i].PublishedAt,
			nil, //TODO あとでテーブルにからむつくる
			account.ID,
			bookTable[i].StartAt,
			bookTable[i].EndAt,
			bookTable[i].NextBookID,
			bookTable[i].PrevBookID,
			descriptionsModel,
			categoriesModel,
			bookTable[i].CreatedAt,
			bookTable[i].UpdatedAt,
		)
		books = append(books, book)
	}
	return &books, nil
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
			bookCategory.Status = pFalse
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
			bookCategory.Status = pTrue
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
	bookTable.PublishedAt = service.NullTime{NullTime:mysql.NullTime{Valid:false }}
	bookTable.StartAt = service.NullTime{NullTime:mysql.NullTime{Valid:false }}
	bookTable.EndAt = service.NullTime{NullTime:mysql.NullTime{Valid:false }}
	bookTable.NextBookID = book.NextBookID
	bookTable.PrevBookID = book.PrevBookID

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
	authorModel := &model.Author{}
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
		authorModel = nil
	}


	err = r.DB.Joins("JOIN books_categories ON books_categories.category_id = categories.id").
		Where("book_id = ?", bookTable[0].ID).
		Where("status = ?", pTrue).
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
	if len(descriptionsTable) != 0 {
		for i := range descriptionsTable {
			description := model.Description{}
			description.Fill(
				descriptionsTable[i].ID,
				descriptionsTable[i].BookId,
				descriptionsTable[i].Description,
				descriptionsTable[i].CreatedAt,
				descriptionsTable[i].UpdatedAt,
			)
			descriptionsModel = append(
				descriptionsModel,
				description,
			)
		}
	}

	book := model.Book{}

	book.Fill(
		bookTable[0].ID,
		bookTable[0].Title,
		authorModel,
		bookTable[0].PublishedAt,
		nil, //TODO あとでテーブルにからむつくる
		account.ID,
		bookTable[0].StartAt,
		bookTable[0].EndAt,
		bookTable[0].NextBookID,
		bookTable[0].PrevBookID,
		descriptionsModel,
		categoriesModel,
		bookTable[0].CreatedAt,
		bookTable[0].UpdatedAt,
	)

	return &book, nil
}

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
	bookTable[0].StartAt = service.NullTime{NullTime:mysql.NullTime{Valid:false }}
	bookTable[0].EndAt = service.NullTime{NullTime:mysql.NullTime{Valid:false }}
	bookTable[0].NextBookID = book.NextBookID
	bookTable[0].PrevBookID = book.PrevBookID

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

func (r *bookRepository) StartReadBook(book model.Book) (*model.Book, error){
	bookTable := tables.Book{}
	bookTable.BindFromModel(book)
	bookTable.StartAt = service.NullTime{NullTime:mysql.NullTime{ Time:time.Now(), Valid:true }}
	bookTable.EndAt = service.NullTime{NullTime:mysql.NullTime{ Valid:false }}
	err := r.DB.Save(&bookTable).Error
	if err != nil {
		return nil, err
	}

	book.StartAt = bookTable.StartAt
	book.EndAt = bookTable.EndAt
	return &book, nil
}
func (r *bookRepository) EndReadBook(book model.Book) (*model.Book, error){
	bookTable := tables.Book{}
	bookTable.BindFromModel(book)
	bookTable.EndAt = service.NullTime{NullTime:mysql.NullTime{ Time:time.Now(), Valid:true }}
	err := r.DB.Save(&bookTable).Error
	if err != nil {
		return nil, err
	}

	book.StartAt = bookTable.StartAt
	book.EndAt = bookTable.EndAt
	return &book, nil
}