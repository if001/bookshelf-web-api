package tables

type BookCategory struct {
	BaseModel
	BookID      int64
	CategoryID  int64
	// Status      int8 `gorm:"default:1" sql:"default:1"` // todo booleanだとうまく更新されないのでとりあえずtinyintでしのぐ
	Status      *bool `gorm:"default:true" sql:"default:true"` // todo booleanだとうまく更新されないのでとりあえずtinyintでしのぐ
}

func (BookCategory) TableName() string {
	return "books_categories"
}


