package tables

import (
	"bookshelf-web-api/domain/model"
	"bookshelf-web-api/domain/service"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Book struct {
	BaseModel
	AccountID   int64
	Title       string `gorm:"type:varchar(40);"`
	AuthorID    sql.NullInt64
	StartAt     service.NullTime
	EndAt       service.NullTime
	PublishedAt service.NullTime
	NextBookID  service.NullInt64
	PrevBookID  service.NullInt64
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
	b.NextBookID = book.NextBookID
	b.PrevBookID = book.PrevBookID
}

type Book2 struct {
	ID                   int64
	AccountID            int64
	Title                string    `gorm:"type:varchar(40);"`
	AuthorID             sql.NullInt64
	StartAt              service.NullTime
	EndAt                service.NullTime
	PublishedAt          service.NullTime
	NextBookID           service.NullInt64
	PrevBookID           service.NullInt64
	CreatedAt            time.Time `gorm;"default:CURRENT_TIMESTAMP" sql:"not null;type:date"`
	UpdatedAt            time.Time `gorm;"default:CURRENT_TIMESTAMP" sql:"not null;type:date"`
	AuthorName           string
	CategoryName         string
	Description          string
	DescriptionCreatedAt time.Time
	Author               Author
}

type SimpleBook struct {
	BaseModel
	AccountID   int64
	Title       string `gorm:"type:varchar(40);"`
	AuthorName  sql.NullString
}