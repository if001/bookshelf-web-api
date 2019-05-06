package model

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
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

type BookRequest struct {
	Title string
	Author string
	Categories []string // default null
	PrevBookId int64 // default 0
	NextBookId int64 // default 0
}