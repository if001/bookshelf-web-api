package model

type Description struct {
	BaseModel
	BookId int64
	Description       string `gorm:"type:varchar(150);"`
}

func (Description) TableName() string {
	return "books_description"
}

