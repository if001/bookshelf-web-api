package tables

type BookCategory struct {
	BaseModel
	BookID     int64
	CategoryID int64
	Status *bool `gorm:"default:true" sql:"default:true"`
}

func (BookCategory) TableName() string {
	return "books_categories"
}
