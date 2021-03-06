package model

import (
	"bookshelf-web-api/domain/service"
	"time"
)

type Book struct {
	Base
	Name         string
	AccountId    int64
	Author       *Author
	PublishedAt  service.NullTime
	Publisher    *Publisher
	StartAt      service.NullTime
	EndAt        service.NullTime
	NextBookID   service.NullInt64
	PrevBookID   service.NullInt64
	Descriptions []Description
	Categories   []Category
}

func (b *Book) GetReadState() ReadState {
	if b.StartAt.Valid && b.EndAt.Valid {
		return ReadValue
	} else if b.StartAt.Valid && !b.EndAt.Valid {
		return ReadingValue
	} else if !b.StartAt.Valid && !b.EndAt.Valid {
		return NotReadValue
	} else {
		return NotReadValue
	}
}

func (b *Book) Fill(id int64, name string, author *Author,
	publishAt service.NullTime, publisher *Publisher,
	accountId int64, startAt service.NullTime, endAt service.NullTime,
	nextBookId service.NullInt64, prevBookId service.NullInt64, descriptions []Description, categories []Category,
	createdAt time.Time, updatedAt time.Time) {
	b.ID = id
	b.Name = name
	b.Author = author
	b.AccountId = accountId
	b.Publisher = nil
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

type Publisher struct {
	Base
	Name string
}
