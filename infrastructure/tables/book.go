package tables

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Book struct {
	BaseModel
	AccountID   int64
	Title       string `gorm:"type:varchar(40);"`
	AuthorID    sql.NullInt64
	StartAt     service.NullTime
	EndAt       service.NullTime
	PublishedAt service.NullTime
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
	b.PublishedAt = service.NullTime{NullTime: mysql.NullTime{Valid: false}}
	b.StartAt = service.NullTime{NullTime: mysql.NullTime{Valid: false}}
	b.EndAt = service.NullTime{NullTime: mysql.NullTime{Valid: false}}
	b.NextBookID = sql.NullInt64{Int64: book.NextBookID, Valid: book.NextBookID != 0}
	b.PrevBookID = sql.NullInt64{Int64: book.PrevBookID, Valid: book.PrevBookID != 0}
}
