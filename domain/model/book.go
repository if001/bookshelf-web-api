package model

import (
	"time"
	"github.com/go-sql-driver/mysql"
)

type Book struct {
	Base
	Name string
	Author *Author
	PublishedAt mysql.NullTime
	Publisher string // todo 構造体にする
	AccountId int64
	StartAt mysql.NullTime
	EndAt mysql.NullTime
	NextBookID  int64
	PrevBookID  int64
	Descriptions []Description
	Categories []Category
}
func (b *Book) GetReadState() ReadState {
	if b.StartAt.Valid && b.EndAt.Valid {
		return &read{}
	} else if b.StartAt.Valid && !b.EndAt.Valid {
		return &reading{}
	} else if !b.StartAt.Valid && !b.EndAt.Valid {
		return &notRead{}
	} else {
		return nil
	}
}
func (b *Book) Fill(id int64, name string, author *Author,
	publishAt mysql.NullTime, publisher string,
	accountId int64, startAt mysql.NullTime, endAt mysql.NullTime,
	nextBookId int64, prevBookId int64, descriptions []Description, categories []Category,
	createdAt time.Time, updatedAt time.Time) {
	b.ID = id
	b.Name = name
	b.Author = author
	b.PublishedAt = publishAt
	b.StartAt = startAt
	b.EndAt = endAt
	b.Publisher = publisher
	b.NextBookID = nextBookId
	b.PrevBookID = prevBookId
	b.Descriptions = descriptions
	b.Categories = categories
	b.CreatedAt = createdAt
	b.UpdatedAt = updatedAt
}


type Category struct {
	Base
	Name string
}
func (a *Category) Fill(id int64, name string, createdAt time.Time, updatedAt time.Time) {
	a.ID = id
	a.Name = name
	a.CreatedAt = createdAt
	a.UpdatedAt = updatedAt
}

type Author struct {
	Base
	Name string
}
func (a *Author) Fill(id int64, name string, createdAt time.Time, updatedAt time.Time) {
	a.ID = id
	a.Name = name
	a.CreatedAt = createdAt
	a.UpdatedAt = updatedAt
}

type Description struct {
	Base
	BookId  int64
	Content string
}
func (a *Description) Fill(id int64, bookId int64, content string, createdAt time.Time, updatedAt time.Time) {
	a.ID = id
	a.Content = content
	a.BookId = bookId
	a.CreatedAt = createdAt
	a.UpdatedAt = updatedAt
}