package tables

type BookCategory struct {
	BaseModel
	BookID      int64
	CategoryID  int64
	Categories  []Category
}

func (BookCategory) TableName() string {
	return "books_categories"
}


