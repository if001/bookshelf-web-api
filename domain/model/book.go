package model

import "time"

type Book struct {
	BaseModel
	AccountID   int64
	Title       string `gorm:"type:varchar(40);"`
	AuthorId    int64
	StartAt     time.Time
	EndAt       time.Time
	PublishedAt time.Time
	NextBookId  int64
	PrevBookId  int64
}

func (Book) TableName() string {
	return "books"
}
