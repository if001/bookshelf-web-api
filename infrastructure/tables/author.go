package tables

type Author struct {
	BaseModel
	Name string `gorm:"type:varchar(30);"`
}

func (Author) TableName() string {
	return "author"
}
