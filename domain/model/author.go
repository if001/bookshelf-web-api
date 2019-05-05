package model


type Author struct {
	BaseModel
	Author string `gorm:"type:varchar(30);"`
}

func (Author) TableName() string {
	return "author"
}

