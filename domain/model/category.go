package model

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(45);"`
}

func (Category) TableName() string {
	return "categories"
}
