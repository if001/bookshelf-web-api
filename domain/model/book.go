package model

import "time"

type Book struct {
	BaseModel
	AccountID   int64
	Title       string `gorm:"type:varchar(40);"`
	AuthorID    int64
	StartAt     time.Time
	EndAt       time.Time
	PublishedAt time.Time
	NextBookID  int64
	PrevBookID  int64
	Author      Author `gorm:"foreignkey:AuthorID"`
}

func (Book) TableName() string {
	return "books"
}
