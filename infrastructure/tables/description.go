package tables

type Description struct {
	BaseModel
	BookId int64
	Description       string `gorm:"type:varchar(150);"`
}

func (Description) TableName() string {
	return "description"
}

