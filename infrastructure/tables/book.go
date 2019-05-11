package tables

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"bookshelf-web-api/domain/model"
)

type Book struct {
	BaseModel
	AccountID   int64
	Title       string `gorm:"type:varchar(40);"`
	AuthorID    sql.NullInt64
	StartAt     mysql.NullTime
	EndAt       mysql.NullTime
	PublishedAt mysql.NullTime
	NextBookID  sql.NullInt64
	PrevBookID  sql.NullInt64
	Author      Author `gorm:"foreignkey:AuthorID"`
	Categories  []Category
	Description []Description
}

func (Book) TableName() string {
	return "books"
}

func (b *Book) BindFromModel(book model.Book) {
	b.ID = book.ID
	b.Title = book.Name
	b.AccountID = book.AccountId
	b.AuthorID = sql.NullInt64{Int64: book.Author.ID, Valid: book.Author != nil}
	b.PublishedAt = mysql.NullTime{Valid:false }
	b.StartAt = mysql.NullTime{Valid:false}
	b.EndAt = mysql.NullTime{ Valid:false}
	b.NextBookID = sql.NullInt64{Int64:book.NextBookID, Valid:book.NextBookID != 0}
	b.PrevBookID = sql.NullInt64{Int64:book.PrevBookID, Valid:book.PrevBookID != 0}
}