package model

import (
	"time"
	"github.com/go-sql-driver/mysql"
)

type Book struct {
	Base
	BookBaseInfo
	BookInfo
}

type BookBaseInfo struct {
	Name string
	Author *Author
	PublishedAt time.Time
	Publisher string
}

type BookInfo struct {
	AccountId string
	StartAt mysql.NullTime
	EndAt mysql.NullTime
	NextBookID  int64
	PrevBookID  int64
	Descriptions []Description
	Categories []Category
}
func (b *BookInfo) GetReadState() ReadState {
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
	Content string
}
func (a *Description) Fill(id int64, content string, createdAt time.Time, updatedAt time.Time) {
	a.ID = id
	a.Content = content
	a.CreatedAt = createdAt
	a.UpdatedAt = updatedAt
}